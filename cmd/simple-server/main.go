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

	"github.com/rafaelfino/metrics"
)

//Structs:
type requestPar struct {
	When   time.Time
	Key    string
	Values []string
}

type Config struct {
	Paths      []ItemFile
	MetricPath string
}

type ItemFile struct {
	HttpPath string
	FilePath string
	Static   bool
	Data     string
}

var config Config
var cache map[string]ItemFile
var metricProcessor *metrics.Processor

//Main
func main() {
	if len(os.Args) < 3 {
		fmt.Println("You need to send the file path and port to serve ")
		return
	}

	metricProcessor = metrics.NewMetricProcessor(time.Second*5, logExport)

	loadConfig(os.Args[1])

	for p, _ := range cache {
		http.HandleFunc(p, handle)
	}

	http.HandleFunc("/refresh", handleRefresh)
	http.HandleFunc("/echo", handleEcho)

	log.Fatal(http.ListenAndServe(os.Args[2], nil))
}

//Handlers
func handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	path := getPath(r.RequestURI)

	if item, found := cache[path]; found {
		if item.Static {
			fmt.Fprintln(w, item.Data)
		} else {
			raw, err := ioutil.ReadFile(item.FilePath)

			if err != nil {
				panic(err)
			}

			fmt.Fprintln(w, string(raw))
		}

		log.Printf(`[%s] request returned in %s`, path, time.Since(start))
		metricProcessor.Send(metrics.NewMetric(path, metrics.CounterType, nil, float64(time.Since(start).Nanoseconds())))
	} else {
		if path != `/favicon.ico` {
			log.Printf(`data not found for uri %s`, path)
		}
	}

	metricProcessor.Send(metrics.NewMetric("http-request", metrics.CounterType, nil, float64(time.Since(start).Nanoseconds())))
}

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

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	path := getPath(r.RequestURI)

	loadCache()

	log.Printf(`[%s] reload cache in %s`, path, time.Since(start))
	metricProcessor.Send(metrics.NewMetric("http-request.echo", metrics.CounterType, nil, float64(time.Since(start).Nanoseconds())))
}

//Internals
func loadConfig(configPath string) {
	cache = make(map[string]ItemFile)

	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		panic(err)
	}

	loadCache()
}

func loadCache() {
	for _, item := range config.Paths {
		if item.Static {
			raw, err := ioutil.ReadFile(item.FilePath)

			if err != nil {
				panic(err)
			}

			item.Data = string(raw)
		}

		cache[item.HttpPath] = item
	}
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
		log.Printf("Metrics: %s\n", string(raw))
		ioutil.WriteFile(config.MetricPath, raw, 777)
	}

	return nil
}
