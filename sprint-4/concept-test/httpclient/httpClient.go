package main

import (
	"360-RedDragons/sprint-4/concept-test/quizjson"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//send GET request to localhost:9000
	resp, err := http.Get(`http://127.0.0.1:9000`)
	if err != nil {
		log.Fatal(`Error: couldn't GET from localhost:9000`)
	}

	//read GET response
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(`Error: couldn't read GET response`)
	}

	//convert to Questions struct
	q := quizjson.ToQuestionSet(msg)
	for _, question := range q.Questions {
		fmt.Println(question)
	}

	resp.Body.Close()
}
