package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"learnops/internal/chat"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	//read config
	config, err := readConfig()

	if err != nil {
		log.Panicf("fail to start: %s", err.Error())
		panic(err)
	}

	//init logger
	f, err := os.OpenFile(config.AppConfig.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		panic(err)
	}

	defer f.Close()
	log.SetOutput(f)

	log.Printf("Config: %v\n", config)

	//init ui manager
	input, list := initUI(config)

	//connect on nats-stream
	conn := chat.NewConn(config.ConnConfig, config.AppConfig.Database)
	//subscribe channel
	ch, err := conn.Subscribe(config.Subject, config.Nickname)

	if err != nil {
		log.Fatalf("fail to subscribe channel: %s\n", err)
	}

	//handle with users and messages
	r := chat.NewRoom(config.Subject, ch)

	//handle with defer resources
	defer func(r *chat.Room, c *chat.Connection) {
		log.Println("Closing connections and elements")

		ui.Close()
		r.Close()
		conn.Close()

		log.Println("Stop!")
	}(r, conn)

	//handle receive messages
	go receive(r.Output, list, config)

	//handle term events and hold execution
	termEventHandler(r, conn, input, list)
}

func readConfig() (*chat.ClientConfig, error) {
	var raw []byte
	var err error

	if raw, err = ioutil.ReadFile(os.Args[1]); err != nil {
		return nil, fmt.Errorf("fail to open config file %s: %s", os.Args[1], err)
	}

	var cfg chat.ClientConfig

	err = json.Unmarshal(raw, &cfg)

	if err != nil {
		return nil, fmt.Errorf("fail to read config file [%s]: %s", os.Args[1], err)
	}

	if cfg.AppConfig == nil {
		cfg.AppConfig = &chat.ChatConfig{}
	}

	if len(cfg.AppConfig.Database) == 0 {
		cfg.AppConfig.Database = os.Args[0] + "." + cfg.Nickname + ".db"
	}

	if len(cfg.AppConfig.LogFile) == 0 {
		cfg.AppConfig.LogFile = os.Args[0] + "." + cfg.Nickname + ".log"
	}

	if len(cfg.ServerURI) == 0 {
		return nil, fmt.Errorf("server-uri not found on %s", os.Args[1])
	}

	if len(cfg.Subject) == 0 {
		return nil, fmt.Errorf("subject not found on %s", os.Args[1])
	}

	if len(cfg.Nickname) == 0 {
		return nil, fmt.Errorf("nickname not found on %s", os.Args[1])
	}

	connRead, err := Get(cfg.ServerURI + `/config`)

	if err != nil {
		return nil, fmt.Errorf("fail to connect on %s: %s", cfg.ServerURI, err)
	}

	var connCfg chat.ConnectionConfig

	err = json.Unmarshal(connRead, &connCfg)

	if err != nil {
		return nil, fmt.Errorf("fail to read config data [%s]: %s", string(connRead), err)
	}

	cfg.ConnConfig = &connCfg

	return &cfg, err
}

func initUI(cfg *chat.ClientConfig) (*widgets.Paragraph, *widgets.List) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v\n", err)
	}

	//creates ui elements
	input := widgets.NewParagraph()
	input.Title = "[" + cfg.Nickname + "]"
	input.BorderStyle = ui.NewStyle(
		ui.ColorYellow,
	)

	input.TitleStyle = ui.NewStyle(
		ui.ColorYellow,
		ui.ColorClear,
		ui.ModifierBold,
	)

	list := widgets.NewList()
	list.Rows = []string{}
	list.Title = "[" + cfg.Subject + "]"
	list.BorderStyle = ui.NewStyle(
		ui.ColorBlue,
	)
	list.TitleStyle = ui.NewStyle(
		ui.ColorBlue,
		ui.ColorClear,
		ui.ModifierBold,
	)

	w, h := ui.TerminalDimensions()
	drawUI(input, list, w, h)

	return input, list
}

func termEventHandler(r *chat.Room, conn *chat.Connection, input *widgets.Paragraph, list *widgets.List) {
	var err error

	for e := range ui.PollEvents() {
		switch e.ID {
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			drawUI(input, list, payload.Width, payload.Height)
		case "<Escape>":
			log.Println("Quit requested")
			return
		case "<Enter>":
			if len(input.Text) == 0 {
				break
			}
			log.Printf("Send message: %s\n", input.Text)
			err = conn.SendMessage(input.Text)

			if err != nil {
				log.Printf("send error: %s\n", err)
			}

			input.Text = ""
		case "<Backspace>", "<C-<Backspace>>":
			if len(input.Text) > 1 {
				input.Text = input.Text[0 : len(input.Text)-1]
			} else {
				input.Text = ""
			}
		case "<Delete>":
			if len(input.Text) > 1 {
				input.Text = input.Text[1:len(input.Text)]
			} else {
				input.Text = ""
			}
		case "<C-c>":
			input.Text = ""
		case "<C-l>":
			r.ListUsers()
		case "<C-r>":
			channels, err := conn.GetChannels()
			if err != nil {
				log.Printf("fail to try get all channels: %s", err)
			}
			r.ListChannels(channels)
		case "<Space>":
			input.Text += " "
		case "<Tab>":
			input.Text += "\t"
		case "<PageUp>":
			list.ScrollTop()
			ui.Render(list)
		case "<PageDown>":
			list.ScrollBottom()
			ui.Render(list)
		case "<Up>":
			list.ScrollUp()
			ui.Render(list)
		case "<Down>":
			list.ScrollDown()
			ui.Render(list)
		default:
			if e.Type == ui.KeyboardEvent {
				if v, found := chat.TbKeys[e.ID]; found {
					input.Text += v
				} else {
					input.Text += e.ID
				}
			}
		}

		ui.Render(input)
	}
}

func receive(ch chan string, list *widgets.List, cfg *chat.ClientConfig) {
	for m := range ch {
		list.Rows = append(list.Rows, m)
		if len(list.Rows) > cfg.MsgBoxBufferSize {
			list.Rows = list.Rows[len(list.Rows)-cfg.MsgBoxBufferSize:]
		}

		list.ScrollBottom()
		ui.Render(list)

		log.Printf("received message: %s\n", m)
	}
}

func drawUI(input *widgets.Paragraph, list *widgets.List, w int, h int) {
	input.SetRect(0, h-3, w, h)
	list.SetRect(0, 0, w, h-3)

	ui.Render(input)
	ui.Render(list)
}

//http tools
func Get(uri string) ([]byte, error) {
	resp, err := http.Get(uri)

	if err != nil {
		log.Printf("fail to execute http get from %s: %s", uri, err)
		return nil, err
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("fail to read http get body from %s: %s", uri, err)
	}

	return ret, err
}
