package chat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ConnectionConfig struct {
	NatsURI       string `json:"nats-uri"`
	ClusterID     string `json:"cluster-id"`
	MonitoringURI string `json:"monitoring-uri"`
}

type ChatConfig struct {
	Database string `json:"database-file,omitempty"`
	LogFile  string `json:"log-file,omitempty"`
}

type ServerConfig struct {
	APIAddress       string            `json:"api-address"`
	ConnectionConfig *ConnectionConfig `json:"connection"`
	AppConfig        *ChatConfig       `json:"app,omitempty"`
}

type ClientConfig struct {
	ScreenSize       int `json:"screen-size,omitempty"`
	MsgBoxSize       int `json:"msg-box-size,omitempty"`
	MsgBoxBufferSize int `json:"msg-box-buffer-size,omitempty"`

	ConnConfig *ConnectionConfig `json:"connection,omitempty"`
	AppConfig  *ChatConfig       `json:"app,omitempty"`

	ServerURI string `json:"server-uri"`

	Subject  string `json:"subject"`
	Nickname string `json:"nickname"`
}

type Message struct {
	When    time.Time   `json:"when"`
	Type    MessageType `json:"type"`
	Subject string      `json:"subject"`
	From    string      `json:"from"`
	Data    string      `json:"data"`
	ID      uint64      `json:"id"`
}

func (m *Message) ToJson() string {
	raw, err := json.Marshal(m)

	if err == nil {
		return string(raw)
	}

	return fmt.Sprint(m)
}

type MessageType int

const (
	Enter MessageType = iota
	UserMessage
	Leave
	KeepAlive
	Kick
	ListUsers
)

var TbKeys = map[string]string{
	"<F1>":        "",
	"<F2>":        "",
	"<F3>":        "",
	"<F4>":        "",
	"<F5>":        "",
	"<F6>":        "",
	"<F7>":        "",
	"<F8>":        "",
	"<F9>":        "",
	"<F10>":       "",
	"<F11>":       "",
	"<F12>":       "",
	"<Insert>":    "",
	"<Home>":      "",
	"<End>":       "",
	"<PageUp>":    "",
	"<PageDown>":  "",
	"<Up>":        "",
	"<Down>":      "",
	"<Left>":      "",
	"<Right>":     "",
	"<C-<Space>>": "",
	"<C-a>":       "",
	"<C-b>":       "",
	"<C-d>":       "",
	"<C-e>":       "",
	"<C-f>":       "",
	"<C-g>":       "",
	"<C-j>":       "",
	"<C-k>":       "",
	"<C-n>":       "",
	"<C-o>":       "",
	"<C-p>":       "",
	"<C-q>":       "",
	"<C-r>":       "",
	"<C-s>":       "",
	"<C-t>":       "",
	"<C-u>":       "",
	"<C-v>":       "",
	"<C-w>":       "",
	"<C-x>":       "",
	"<C-y>":       "",
	"<C-z>":       "",
	"<C-4>":       "",
	"<C-5>":       "",
	"<C-6>":       "",
	"<C-7>":       "",
}

type NatsChannels struct {
	ClusterID string   `json:"cluster_id"`
	ServerID  string   `json:"server_id"`
	Offset    int64    `json:"offset"`
	Limit     int64    `json:"limit"`
	Count     int64    `json:"count"`
	Total     int64    `json:"total"`
	Names     []string `json:"names"`
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
