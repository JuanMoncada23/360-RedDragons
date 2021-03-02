package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9000")
	if err != nil {
		log.Fatal("Error: couldn't GET")
	}
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error: couldn't READ")
	}
	fmt.Println(string(msg))
}
