package chat

import (
	"fmt"
	"log"
	"strings"

	"time"
)

type Room struct {
	users    map[string]time.Time
	received chan *Message
	messages []*Message
	holder   chan bool
	subject  string
	Output   chan string
}

//New Creates a new chat room and handle messages
func NewRoom(subject string, ch chan *Message) *Room {
	room := &Room{
		received: ch,
		users:    map[string]time.Time{},
		messages: []*Message{},
		subject:  subject,
		Output:   make(chan string, 2000),
		holder:   make(chan bool),
	}

	go room.processMessages()
	go room.autoKick()

	return room
}

func (r *Room) Close() {
	r.holder <- true
}

func (r *Room) ListUsers() {
	ret := []string{}

	for u := range r.users {
		ret = append(ret, u)
	}

	r.Output <- fmt.Sprintf("[%s] Active users: %s", time.Now().Format("2006-01-02 15:04:05"), strings.Join(ret, ", "))
}

func (r *Room) ListChannels(channels []string) {
	r.Output <- fmt.Sprintf("[%s] Active channels: %s", time.Now().Format("2006-01-02 15:04:05"), strings.Join(channels, ", "))
}

func (r *Room) processMessages() {
	for msg := range r.received {
		switch msg.Type {
		case Leave:
			if _, found := r.users[msg.From]; found {
				delete(r.users, msg.From)
			}
			r.receiveMessage(msg)
		case UserMessage, Enter:
			r.receiveMessage(msg)
			r.users[msg.From] = time.Now()
		case KeepAlive:
			log.Printf("Keep alive received from %s\n", msg.From)
			r.users[msg.From] = time.Now()
		}
	}
}

func (r *Room) autoKick() {
	for {
		select {
		case <-time.After(time.Second):
			for u, t := range r.users {
				if time.Since(t) > (time.Minute) {
					r.receiveMessage(&Message{
						When:    time.Now(),
						Data:    u + " is no longer active",
						Subject: r.subject,
						From:    u,
						Type:    Kick,
					})
					delete(r.users, u)
				}
			}
		case <-r.holder:
			return
		}
	}
}

func (r *Room) receiveMessage(msg *Message) {
	switch msg.Type {
	case Enter, Leave, Kick:
		r.Output <- fmt.Sprintf("[%s] %s", msg.When.Format("2006-01-02 15:04:05"), msg.Data)
	default:
		r.Output <- fmt.Sprintf("[%s] %s: %s", msg.When.Format("2006-01-02 15:04:05"), msg.From, msg.Data)
	}
}
