Proof-of-Concept
================

HTML
----
The HTML Server serves data on localhost:9090
* run this in its own terminal window or as a background process

The HTML Client performs a GET request to the HTML Server
* run this in your code editor

JSON Parsing
------------
The JSON Reader parses data from example.json into a struct which can be printed out.

The struct and all of its properties MUST be exported or it will break

Credit
------
https://golangdocs.com/golang-read-json-file
* Example of parsing JSON data from a .json file

tour.golang.org/methods/17
* Example of implementing Stringer interface (allows fmt.Println(Question_Struct))