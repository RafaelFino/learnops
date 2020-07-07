package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"learnops/internal/chat"

	"github.com/rafaelfino/metrics"
)

//Structs:
type requestPar struct {
	When   time.Time
	Key    string
	Values []string
}

var cfg *chat.ServerConfig
var metricProcessor *metrics.Processor
var conn *chat.Connection

func main() {
	//read config
	err := readConfig()

	if err != nil {
		panic(err)
	}

	//init logger
	f, err := os.OpenFile(cfg.AppConfig.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	metricProcessor = metrics.NewMetricProcessor(time.Second*5, logExport)

	if err != nil {
		panic(err)
	}

	defer f.Close()
	log.SetOutput(f)

	log.Printf("Config: %v\n", cfg)

	//connect on nats-stream
	conn = chat.NewConn(cfg.ConnectionConfig, cfg.AppConfig.Database)
	//subscribe channel
	ch, err := conn.SubscribeServer()

	if err != nil {
		log.Fatalf("fail to subscribe channel: %s\n", err)
	}

	//handle with defer resources
	go func(c *chat.Connection) {
		for m := range ch {
			log.Printf("[Received message] %s\n", m.ToJson())
		}

		log.Println("Closing connections and elements")
		conn.Close()

		log.Println("Stop!")
	}(conn)

	http.HandleFunc("/channels", handleGetChannels)
	http.HandleFunc("/config", handleGetConfig)
	http.HandleFunc("/echo", handleEcho)

	log.Fatal(http.ListenAndServe(cfg.APIAddress, nil))
}

func readConfig() error {
	var raw []byte
	var err error

	if raw, err = ioutil.ReadFile(os.Args[1]); err != nil {
		return err
	}

	var c chat.ServerConfig

	err = json.Unmarshal(raw, &c)

	if c.AppConfig == nil {
		c.AppConfig = &chat.ChatConfig{}
	}

	if len(c.AppConfig.Database) == 0 {
		c.AppConfig.Database = os.Args[0] + ".server.db"
	}

	if len(c.AppConfig.LogFile) == 0 {
		c.AppConfig.LogFile = os.Args[0] + ".server.log"
	}

	cfg = &c

	return err
}

//handlers
func handleEcho(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	path := getPath(r.RequestURI)

	data := make([]requestPar, 0)
	for key, value := range r.URL.Query() {
		data = append(data, requestPar{When: time.Now(), Key: key, Values: value})
	}
	raw, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("<html><body><b>Echo error:</b><br>%s</body></html>", err))
	} else {
		fmt.Fprintln(w, string(raw))
	}

	log.Printf(`[%s] request returned in %s`, path, time.Since(start))
	metricProcessor.Send(metrics.NewMetric("http-request.echo", metrics.CounterType, nil, float64(time.Since(start).Nanoseconds())))
}

func handleGetConfig(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	path := getPath(r.RequestURI)

	raw, err := json.MarshalIndent(cfg.ConnectionConfig, "", "\t")

	if err != nil {
		log.Printf("fail to marshal config: %s\n", err)
		return
	}

	fmt.Fprintln(w, string(raw))

	metricProcessor.Send(metrics.NewMetric("http-request.get-config", metrics.CounterType, nil, float64(time.Since(start).Nanoseconds())))

	log.Printf(`[http-request] [%s] request returned in %s`, path, time.Since(start))
}

func handleGetChannels(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	path := getPath(r.RequestURI)

	channels, err := conn.GetChannels()

	if err != nil {
		log.Printf("fail to get nats channels: %s\n", err)
		return
	}

	raw, err := json.MarshalIndent(channels, "", "\t")

	if err != nil {
		log.Printf("fail to marshal channels: %s\n", err)
		return
	}

	fmt.Fprintln(w, string(raw))

	metricProcessor.Send(metrics.NewMetric("http-request.get-channels", metrics.CounterType, nil, float64(time.Since(start).Nanoseconds())))
	log.Printf(`[http-request] [%s] request returned in %s`, path, time.Since(start))
}

func getPath(path string) string {
	idx := strings.Index(path, `?`)

	if idx > 0 {
		return string(path[0:idx])
	}

	return path
}

//Export metrics
func logExport(data *metrics.MetricData) error {
	if len(data.Metrics) == 0 && len(data.Series) == 0 {
		return nil
	}

	raw, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		log.Printf("fail to marshal metrics: %s\n", err)
	} else {
		log.Printf("[Metrics] %s", string(raw))
	}

	return nil
}
