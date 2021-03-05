package main

import (
	"360-RedDragons/sprint-4/concept-test/quizjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func info(res http.ResponseWriter, req *http.Request) {
	//find JSON question banks to use
	files, err := ioutil.ReadDir(`..`)
	if err != nil {
		log.Fatal(`Error: couldn't read JSON directory`)
	}

	//find and collect all relevant JSON question banks
	b := quizjson.Banks{}
	for _, fi := range files {
		if strings.Contains(fi.Name(), `.json`) {
			filename := `..` + string(os.PathSeparator) + fi.Name()
			/* jsonFiles = append(jsonFiles, fileName) */
			e := quizjson.FromFile(filename)
			b.Banks = append(b.Banks, quizjson.Bank{e.Name, e.ID, len(e.Entries)})
		}
	}

	//return in JSON format
	bytes := b.ToJSON()
	io.WriteString(res, string(bytes))
}

func hello(res http.ResponseWriter, req *http.Request) {
	//read GET request

	//retrieve entries
	e := quizjson.FromFile(`../bank-1.json`)

	//convert to questions
	q := quizjson.Questions{}
	for _, entry := range e.Entries {
		q.Questions = append(q.Questions, entry.Question)
	}

	//return in JSON format
	content := q.ToJSON()
	io.WriteString(res, string(content))
}

func main() {
	http.HandleFunc(`/`, hello)
	http.HandleFunc(`/info`, info)
	http.ListenAndServe(`:9000`, nil)
}
