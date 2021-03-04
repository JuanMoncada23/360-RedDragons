package main

import (
	"360-RedDragons/sprint-4/QuizMaster/quizjson"
	"encoding/json"
	"fmt"
)

func main() {
	entries := quizjson.FromFile(`./example.json`)
	bytes, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(bytes))
}
