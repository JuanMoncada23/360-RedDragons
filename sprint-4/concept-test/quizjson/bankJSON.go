package quizjson

import (
	"encoding/json"
	"log"
)

//Banks struct represents a set of JSON-based quiz files
type Banks struct {
	Banks []Bank `json:"banks"`
}

//Bank struct represents one JSON-based quiz file
type Bank struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	QuestionCt int    `json:"questionCt"`
}

//ToJSON method returns Banks struct as JSON []byte
func (b Banks) ToJSON() []byte {
	bytes, err := json.Marshal(b)
	if err != nil {
		log.Fatal(`Error: couldn't convert Banks to JSON`)
	}
	return bytes
}
