package main

import (
	"log"
	"net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    log.Print("starting server on port:8080")

    log.Fatal(http.ListenAndServe(":8080", mux))
}
