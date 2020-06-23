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
)

//Structs:
type requestPar struct {
	When   time.Time
	Key    string
	Values []string
}

type Config struct {
	Paths map[string]string
}

var cache map[string]string

//Main
func main() {
	if len(os.Args) < 3 {
		fmt.Println("You need to send the file path and port to serve ")
		return
	}

	loadConfig(os.Args[1])

	for p, _ := range cache {
		http.HandleFunc(p, handle)
	}

	http.HandleFunc("/echo", handleEcho)

	log.Fatal(http.ListenAndServe(os.Args[2], nil))
}

//Handlers
func handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	path := getPath(r.RequestURI)

	if ret, found := cache[path]; found {
		fmt.Fprintln(w, ret)
	} else {
		log.Printf(`data not found for uri %s`, path)
	}

	log.Printf(`[%s] request returned in %s`, path, time.Since(start))
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
}

//Internals
func loadConfig(configPath string) {
	cache = make(map[string]string)

	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = json.Unmarshal(data, &cfg)

	if err != nil {
		panic(err)
	}

	for path, file := range cfg.Paths {
		raw, err := ioutil.ReadFile(file)

		if err != nil {
			panic(err)
		}

		cache[path] = string(raw)
	}
}

func getPath(path string) string {
	idx := strings.Index(path, `?`)

	if idx > 0 {
		return string(path[0:idx])
	}

	return path
}
