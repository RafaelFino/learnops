package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var url = `http://loripsum.net/api`

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("you need to inform output file path"))
	}

	log.Printf("openning file: %s\n", os.Args[1])

	f, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	qty := 5

	for i := 1; i <= qty; i++ {
		log.Printf("[%d/%d] executing request...\n", i, qty)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		log.Println("writing file...")
		if _, err = f.WriteString(string(body) + "\n"); err != nil {
			panic(err)
		}
	}
	log.Printf("stop!")
}
