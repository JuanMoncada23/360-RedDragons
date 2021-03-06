package quizjson

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

//Questions struct represents a set of JSON-based quiz questions
type Questions struct {
	Questions []Question `json:"questions"`
}

//Question struct represents one JSON-based quiz question
type Question struct {
	Prompt  string   `json:"prompt"`
	Choices []string `json:"choices"`
}

//String method allows printing Question struct as a string
func (q Question) String() string {
	retStr := q.Prompt + "\n"
	for _, i := range q.Choices {
		retStr += i + "\n"
	}
	return fmt.Sprintf(retStr)
}

//ShuffleChoices method randomly shuffles the ordering of choices
func (q *Question) ShuffleChoices() {
	//prep random number generator
	rand.Seed(time.Now().UnixNano())

	//copy q.Choices into temporary slice
	tmpChoice := make([]string, len(q.Choices))
	for i, choice := range q.Choices {
		tmpChoice[i] = choice
	}

	for i := range q.Choices {
		//choose random index into tmpChoice
		choice := rand.Intn(int(time.Now().UnixNano())) % len(tmpChoice)
		q.Choices[i] = tmpChoice[choice]

		//remove value at index `choice` in tmpChoice
		if choice == 0 {
			tmpChoice = tmpChoice[1:]
		} else if choice == len(tmpChoice)-1 {
			tmpChoice = tmpChoice[:len(tmpChoice)-1]
		} else {
			lo := tmpChoice[:choice]
			hi := tmpChoice[choice+1:]
			tmpChoice = lo
			tmpChoice = append(tmpChoice, hi...)
		}
	}
}

//ToJSON method returns Questions struct as JSON []byte
func (q Questions) ToJSON() []byte {
	bytes, err := json.Marshal(q)
	if err != nil {
		log.Fatal(`Error: couldn't convert Questions struct to JSON`)
	}
	return bytes
}

//ToQuestionSet function returns JSON []byte as Questions struct
func ToQuestionSet(bytes []byte) Questions {
	var q Questions
	err := json.Unmarshal(bytes, &q)
	if err != nil {
		log.Fatal(`Error: couldn't convert JSON to Questions struct`, err)
	}
	return q
}
