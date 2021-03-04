package main

import (
	"360-RedDragons/sprint-4/QuizMaster/concept-test/quizjson"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//send GET request to localhost:9000
	resp, err := http.Get("http://127.0.0.1:9000")
	if err != nil {
		//log Issue
		log.Fatal("Issue: couldn't GET from localhost:9000")
	}

	//read GET response from localhost:9000
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log issue
		log.Fatal("Error: couldn't read GET response")
	}

	//convert to struct and print response
	q := quizjson.ToQuestion(msg)
	fmt.Println(q)

	//close response body
	resp.Body.Close()
}
