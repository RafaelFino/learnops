package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var url = `https://economia.awesomeapi.com.br/json/all`

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("you need to inform output file path"))
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	f, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(string(body))

	if err != nil {
		panic(err)
	}
}
