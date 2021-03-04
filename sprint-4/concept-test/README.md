Proof-of-Concept
================

HTTP
----
The HTTP Server serves data on localhost:9090
* run this in its own terminal window or as a background process
* you can visit localhost:9000 in your browser to see the JSON data
* you can visit localhost:9000/info in your browser to see the JSON bank info

The HTTP Client performs a GET request to the HTTP Server
* run this in your code editor

JSON Parsing
------------
There is a JSON Parser for bank entries (includes prompt, choices, and answer).
* this is jsonquiz/entryJSON.go
There is a JSON Parser for question entries (includes prompt and choices).
* this is jsonquiz/questionJSON.go

These are used by the HTTP Server and Client .go files to send, receive, and interpret
data. They are also used by main.go for testing / troubleshooting the various structs
and methods.

JSON structs and all of their properties MUST be exported or it will not work

Credit
------
https://golangdocs.com/golang-read-json-file
* Example of parsing JSON data from a .json file

tour.golang.org/methods/17
* Example of implementing Stringer interface (allows fmt.Println(Question_Struct))