package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	w.Write([]byte("testing..."))
}

func showSnippet(w http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, req)
		return
	}

	fmt.Fprintf(w, "This is snippet with ID %d", id)
	//w.Write([]byte("Show snippet"))
}

func createSnippet(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Get method not allowed!", http.StatusMethodNotAllowed)
		//w.WriteHeader(405)
		//w.Write([]byte("GET method is forbidden!"))
		return
	}

	w.Write([]byte("Create snippet"))
}
