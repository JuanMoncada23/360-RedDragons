package quizjson

import (
	"encoding/json"
	"fmt"
	"log"
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

//ToJSON method returns Questions struct as JSON []byte
func (q Questions) ToJSON() []byte {
	bytes, err := json.Marshal(q)
	if err != nil {
		log.Fatal(`Error: couldn't convert Questions to JSON`)
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
