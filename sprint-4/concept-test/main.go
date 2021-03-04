package main

import (
	"360-RedDragons/sprint-4/concept-test/quizjson"
	"encoding/json"
	"fmt"
)

func main() {
	entries := quizjson.FromFile(`./bank-1.json`)
	bytes, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(bytes))
}
