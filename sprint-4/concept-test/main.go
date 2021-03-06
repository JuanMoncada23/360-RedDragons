package main

import (
	"360-RedDragons/sprint-4/concept-test/quizjson"
	"fmt"
)

func main() {
	e := quizjson.FromFile(`./bank-1.json`)
	q := e.Entries[0].Question
	q.ShuffleChoices()
	fmt.Println(q)
}
