package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `JSON`)
}

func main() {
	http.HandleFunc(`/`, hello)
	http.ListenAndServe(`:9000`, nil)
}
