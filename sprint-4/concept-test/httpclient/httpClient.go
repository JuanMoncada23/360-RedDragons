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
	req := quizjson.ReqJSON{IDs: []string{`1`, `2`}, Count: []int{3, 3}}
	bs := req.ToJSON()
	resp, err := http.Post(`http://127.0.0.1:9000/req`, `application/json`, bytes.NewBuffer(bs))
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
}
