package main

import (
	"360-RedDragons/sprint-4/concept-test/quizjson"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var idFileMap map[string]string = make(map[string]string)

func init() {
	//open parent directory
	files, err := ioutil.ReadDir(`..`)
	if err != nil {
		log.Fatal(`Error: couldn't read JSON directory`)
	}

	//look for JSON question banks
	for _, fi := range files {
		if strings.Contains(fi.Name(), `.json`) {
			//associate each bank id w/ filename in global map
			filename := `..` + string(os.PathSeparator) + fi.Name()
			e := quizjson.FromFile(filename)
			idFileMap[e.ID] = filename
		}
	}
	fmt.Println("done")
}

func info(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	//collect all relevant JSON question banks
	b := quizjson.Banks{}
	for id, file := range idFileMap {
		e := quizjson.FromFile(file)
		b.Banks = append(b.Banks, quizjson.Bank{Name: e.Name, ID: id, Count: len(e.Entries)})
	}

	//return in JSON format
	bytes := b.ToJSON()
	io.WriteString(res, string(bytes))
}

func hello(res http.ResponseWriter, req *http.Request) {
	//read POST request
	defer req.Body.Close()
	msg, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(`Error: couldn't read POST request`)
	}

	//retrieve entries and convert to questions
	e := quizjson.FromFile(`../bank-1.json`)
	q := quizjson.Questions{}
	for _, entry := range e.Entries {
		q.Questions = append(q.Questions, entry.Question)
	}

	//return in JSON format
	msg = q.ToJSON()
	io.WriteString(res, string(msg))

	/* //read GET request
	defer req.Body.Close()

	//retrieve entries
	e := quizjson.FromFile(`../bank-1.json`)

	//convert to questions
	q := quizjson.Questions{}
	for _, entry := range e.Entries {
		q.Questions = append(q.Questions, entry.Question)
	}

	//return in JSON format
	content := q.ToJSON()
	io.WriteString(res, string(content)) */
}

func main() {
	http.HandleFunc(`/`, hello)
	http.HandleFunc(`/info`, info)
	http.ListenAndServe(`:9000`, nil)
}
