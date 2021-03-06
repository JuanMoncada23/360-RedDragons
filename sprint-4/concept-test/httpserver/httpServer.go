package main

import (
	"360-RedDragons/sprint-4/concept-test/quizjson"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
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

	//retrieve entries, convert to questions, shuffle answer choices
	e := quizjson.FromFile(`../bank-1.json`)
	q := quizjson.Questions{}
	for _, entry := range e.Entries {

		//shuffle question choices between 5 and 30 times
		rand.Seed(time.Now().UnixNano())
		shaker := rand.Intn(26) + 5
		for i := 0; i < shaker; i++ {
			time.Sleep(time.Duration(50))
			entry.Question.ShuffleChoices()
		}
		q.Questions = append(q.Questions, entry.Question)
	}

	//return in JSON format
	msg = q.ToJSON()
	io.WriteString(res, string(msg))
}

func main() {
	http.HandleFunc(`/`, hello)
	http.HandleFunc(`/info`, info)
	http.ListenAndServe(`:9000`, nil)
}
