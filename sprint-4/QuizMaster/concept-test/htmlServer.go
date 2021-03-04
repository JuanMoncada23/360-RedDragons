package main

import (
	"360-RedDragons/sprint-4/QuizMaster/concept-test/quizjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	//read GET request
	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		//log issue
		log.Fatal("Issue: couldn't read GET request")
	}

	//retrieve entry
	entry := quizjson.ReadFromFile(`./example.json`)

	//convert to question struct (to transmit over
	//insecure http connection)
	//then convert to JSON []byte for writing
	content = entry.ToQuestion().ToJSON()

	//write JSON response
	io.WriteString(res, string(content))
}

func main() {
	http.HandleFunc(`/`, hello)
	http.ListenAndServe(`:9000`, nil)
}
