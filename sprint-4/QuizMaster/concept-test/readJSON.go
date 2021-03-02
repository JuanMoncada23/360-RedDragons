package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Question struct {
	Prompt    string   `json:"prompt"`
	Responses []string `json:"responses`
	Answer    string   `json:"answer"`
}

func (q Question) String() string {
	return fmt.Sprintf("Prompt: %s\nResponses: %v\nAnswer: %s",
		q.Prompt, q.Responses, q.Answer)
}

func main() {
	bytes, err := ioutil.ReadFile(`./example.json`)
	if err != nil {
		fmt.Println("Issue:", err)
	}

	var q Question
	json.Unmarshal(bytes, &q)

	fmt.Println(q)
}
