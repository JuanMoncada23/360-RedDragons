package main

import (
	"360-RedDragons/sprint-4/concept-test/quizjson"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//prep and send POST request to localhost:9000
	req := quizjson.ReqJSON{[]string{`1`}, []int{3}}
	bs := req.ToJSON()
	resp, err := http.Post(`http://127.0.0.1:9000`, `application/json`, bytes.NewBuffer(bs))
	if err != nil {
		log.Fatal(`Error: couldn't POST from localhost:9000`)
	}
	defer resp.Body.Close()

	//catch response
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(`Error: couldn't read POST response`)
	}

	//interpret response
	q := quizjson.ToQuestionSet(msg)
	for _, question := range q.Questions {
		fmt.Println(question)
	}

	//send GET request to localhost:9000
	/* resp, err := http.Get(`http://127.0.0.1:9000`)
	if err != nil {
		log.Fatal(`Error: couldn't GET from localhost:9000`)
	}
	defer resp.Body.Close()

	//read GET response
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(`Error: couldn't read GET response`)
	}

	//convert to Questions struct
	q := quizjson.ToQuestionSet(msg)
	for _, question := range q.Questions {
		fmt.Println(question)
	} */
}
