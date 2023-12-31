package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(ExampleHandler))
	http.ListenAndServe(":8080", nil)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "This is a get req")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "This is a post req")
	} else {
		io.WriteString(w, "This is a "+r.Method+"req")
	}
}
