package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
)

//Connection internal struct
type Connection struct {
	config         *ConnectionConfig
	subject        string
	nickname       string
	receiveChannel chan *Message
	conn           stan.Conn
	storage        *Storage
}

//New creates a connection
func NewConn(cfg *ConnectionConfig, database string) *Connection {
	ret := &Connection{
		config:         cfg,
		receiveChannel: make(chan *Message, 2000),
	}

	storage, err := NewStorage(database)

	if err != nil {
		log.Fatalf("fail to try open database [%s] on %s", database, err)
	}

	ret.storage = storage

	return ret
}

func (c *Connection) SubscribeServer() (chan *Message, error) {
	channels, err := c.GetChannels()

	if err != nil {
		log.Fatalf("get channels error: %s\n", err)
		return nil, err
	}

	log.Printf("Trying to connect on %s\n", c.config.NatsURI)
	sc, err := stan.Connect(c.config.ClusterID, c.config.ClusterID+"-chat-server", stan.NatsURL(c.config.NatsURI))

	if err != nil {
		log.Fatalf("nats connect error: %s\n", err)
		return nil, err
	}

	c.conn = sc
	log.Printf("Connected on %s\n", c.config.NatsURI)

	for _, channel := range channels {
		if pastMessages, err := c.storage.Read(channel, "server"); err == nil {
			for _, m := range pastMessages {
				c.receiveChannel <- m
			}
		} else {
			log.Fatalf("error to try recovery past messages, err: %s", err)
		}

		log.Printf("Trying to subscribe on %s\n", channel)
		_, err = sc.Subscribe(
			channel,
			c.receiveMessage,
			stan.StartAtTimeDelta(2*time.Hour),
			stan.DurableName(channel),
		)

		if err != nil {
			return nil, err
		}
		log.Printf("Subscribed on %s\n", channel)
	}

	return c.receiveChannel, err
}

//Subscribe subscribe a subject
func (c *Connection) Subscribe(subject string, nickname string) (chan *Message, error) {
	log.Printf("Trying to connect on %s\n", c.config.NatsURI)
	sc, err := stan.Connect(c.config.ClusterID, fmt.Sprintf("%s_%s", subject, nickname), stan.NatsURL(c.config.NatsURI))

	if err != nil {
		return nil, err
	}

	c.conn = sc

	c.subject = subject
	c.nickname = nickname

	if pastMessages, err := c.storage.Read(c.subject, c.nickname); err == nil {
		for _, m := range pastMessages {
			c.receiveChannel <- m
		}
	} else {
		log.Fatalf("error to try recovery past messages, err: %s", err)
	}

	log.Printf("Trying to subscribe on %s\n", c.subject)
	_, err = sc.Subscribe(
		c.subject,
		c.receiveMessage,
		stan.StartAtTimeDelta(2*time.Hour),
		stan.DurableName(c.subject),
	)

	if err != nil {
		return nil, err
	}

	if err = c.sendEnter(); err != nil {
		log.Fatalf("subscribe error: %s\n", err)
		return c.receiveChannel, err
	}

	go c.keepAlive()

	log.Printf("Connected on %s\n", c.config.NatsURI)

	return c.receiveChannel, err
}

//SendMessage Send an user message
func (c *Connection) SendMessage(data string) error {
	raw, err := json.Marshal(Message{
		When:    time.Now(),
		Data:    data,
		Subject: c.subject,
		From:    c.nickname,
		Type:    UserMessage,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.conn.Publish(c.subject, raw)
}

func (c *Connection) sendEnter() error {
	raw, err := json.Marshal(Message{
		When:    time.Now(),
		Data:    c.nickname + " is here!",
		Subject: c.subject,
		From:    c.nickname,
		Type:    Enter,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.conn.Publish(c.subject, raw)
}

func (c *Connection) keepAlive() {
	for c.conn != nil {
		time.Sleep(45 * time.Second)

		raw, err := json.Marshal(Message{
			When:    time.Now(),
			Subject: c.subject,
			From:    c.nickname,
			Type:    KeepAlive,
		})

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Publishing keep alive on %s\n", c.config.NatsURI)
		c.conn.Publish(c.subject, raw)
	}
}

func (c *Connection) sendLeave() error {
	raw, err := json.Marshal(Message{
		When:    time.Now(),
		Data:    c.nickname + " quit",
		Subject: c.subject,
		From:    c.nickname,
		Type:    Leave,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.conn.Publish(c.subject, raw)
}

//Close Close streaming connection
func (c *Connection) Close() {
	log.Printf("Closing nats-stream connection")

	if c.conn != nil {
		if err := c.sendLeave(); err != nil {
			panic(err)
		}

		c.conn.Close()
		c.conn = nil
		close(c.receiveChannel)
	}
}

func (c *Connection) receiveMessage(msg *stan.Msg) {
	var received Message

	if err := json.Unmarshal(msg.Data, &received); err == nil {
		received.ID = msg.Sequence
		c.receiveChannel <- &received

		if received.Type != KeepAlive {
			if err = c.storage.Write(&received); err != nil {
				log.Fatalf("fail to try write message on db, msg: %v, err: %s", received, err)
			}
		}
	}

	msg.Ack()
}

//GetChannels Return all nats channels
func (c *Connection) GetChannels() ([]string, error) {
	response, err := Get(c.config.MonitoringURI + `/streaming/channelsz`)

	var channels NatsChannels

	err = json.Unmarshal(response, &channels)

	if err != nil {
		log.Printf("fail to get nats channels info: %s\n", err)
		return nil, err
	}

	if channels.Names == nil {
		channels.Names = []string{}
	}

	raw, err := json.MarshalIndent(channels, "", "\t")

	log.Printf("channels info: %s\n", string(raw))

	return channels.Names, err
}
