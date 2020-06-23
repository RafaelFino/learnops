package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//Structs:
type requestPar struct {
	When   time.Time
	Key    string
	Values []string
}

var filePath string
var cache string

//Main
func main() {
	if len(os.Args) < 3 {
		fmt.Println("You need to send the file path and port to serve ")
		return
	}

	filePath = os.Args[1]

	http.HandleFunc("/", handle)
	http.HandleFunc("/echo", handleEcho)

	log.Fatal(http.ListenAndServe(os.Args[2], nil))
}

//Handlers
func handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	fmt.Fprintln(w, getFile(filePath))

	log.Printf(`request received on / in %s`, time.Since(start))
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

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

	log.Printf(`request received on /echo in %s`, time.Since(start))
}

//Internals
func getFile(path string) string {
	if len(cache) == 0 {
		data, err := ioutil.ReadFile(path)

		if err != nil {
			panic(err)
		}

		cache = string(data)
	}

	return cache
}
