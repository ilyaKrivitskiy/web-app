package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("testing..."))
}

func showSnippet(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Show snippet"))
}

func createSnippet(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Create snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Start server on http://127.0.0.1:4000")
	errorOfServer := http.ListenAndServe(":4000", mux)
	log.Println(errorOfServer)
}
