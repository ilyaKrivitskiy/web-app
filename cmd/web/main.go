package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Start server on http://127.0.0.1:4000")
	errorOfServer := http.ListenAndServe(":4000", mux)
	log.Fatal(errorOfServer)
}
