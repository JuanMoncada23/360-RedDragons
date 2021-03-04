package quizjson

import (
	"encoding/json"
	"fmt"
	"log"
)

//Question struct for JSON question entries
type Question struct {
	Prompt  string   `json:"prompt"`
	Choices []string `json:"responses"`
}

func (q Question) String() string {
	retStr := q.Prompt + "\n"
	for _, i := range q.Choices {
		retStr += i + "\n"
	}
	return fmt.Sprintf(retStr)
}

/*ToJSON returns Question struct as JSON []byte

var content []byte = q.QuestionToJSON()*/
func (q Question) ToJSON() []byte {
	//pack Question struct to JSON []byte
	content, err := json.Marshal(q)
	if err != nil {
		//log issue
		log.Fatal(`Error: couldn't pack Question struct to JSON`)
	}

	//return []byte
	return content
}

/*QuestionToJSON returns Question params as JSON []byte

var content []byte = QuestionToJSON("prompt", "c1", "c2", ...)*/
func QuestionToJSON(prompt string, choices ...string) []byte {
	//create Question struct to convert
	obj := Question{prompt, choices}

	return obj.ToJSON()
}

/*ToQuestion returns bytes param as Question struct

var q Question = ToQuestion([]byte)*/
func ToQuestion(bytes []byte) Question {
	var q Question
	err := json.Unmarshal(bytes, &q)
	if err != nil {
		//log issue
		log.Fatal(`Error: couldn't pack JSON to Question struct`)
	}
	return q
}
