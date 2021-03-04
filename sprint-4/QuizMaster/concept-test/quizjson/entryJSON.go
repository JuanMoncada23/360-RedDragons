package quizjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//Entry struct for JSON question entries
type Entry struct {
	Prompt  string   `json:"prompt"`
	Answer  string   `json:"answer"`
	Choices []string `json:"responses"`
}

//ToQuestion returns Entry struct as a Question struct
func (e Entry) ToQuestion() Question {
	return Question{e.Prompt, e.Choices}
}

func (e Entry) String() string {
	retStr := e.Prompt + "\n"
	for _, i := range e.Choices {
		retStr += i + "\n"
	}
	retStr += "Answer: " + e.Answer
	return fmt.Sprintf(retStr)
}

/*ToJSON returns Entry struct as JSON []byte

var content []byte = e.EntryToJSON()*/
func (e Entry) ToJSON() []byte {
	//pack Entry struct to JSON []byte
	content, err := json.Marshal(e)
	if err != nil {
		//log issue
		log.Fatal(`Error: couldn't pack Entry struct to JSON`)
	}

	//return []byte
	return content
}

/*EntryToJSON returns Entry params as JSON []byte

var content []byte = ConvertToJSON("prompt", "correct", "c1", "c2", ...)*/
func EntryToJSON(prompt, correct string, choices ...string) []byte {
	//create Entry struct to convert
	obj := Entry{prompt, correct, choices}

	return obj.ToJSON()
}

/*ToEntry returns bytes param as Entry struct

var e Entry = ToEntry([]byte)*/
func ToEntry(bytes []byte) Entry {
	var e Entry
	err := json.Unmarshal(bytes, &e)
	if err != nil {
		//log issue
		log.Fatal(`Error: couldn't pack JSON to Entry struct`)
	}
	return e
}

/*ReadFromFile returns Entry struct

var e Entry = ReadFromJSONFile("./example.json")*/
func ReadFromFile(filepath string) Entry {
	//open and read file
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		//log issue
		log.Fatal(`Error: couldn't open JSON file`)
	}

	//unpack JSON into Entry struct
	e := ToEntry(bytes)

	//return Entry struct
	return e
}

/* READ FROM JSON FILE
entry := ReadFromFile(`./example.json`)
fmt.Println(entry)

WRITE TO Entry Struct
content := EntryToJSON(`hello`, `c`, `a`, `b`, `c`)
fmt.Println(string(content))*/
