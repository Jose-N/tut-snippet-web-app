package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
    addr := flag.String("addr", ":8080", "HTTP network address")
    flag.Parse()

    mux := http.NewServeMux()
    fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("/static/", http.StripPrefix("/static", fileServer))
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    log.Printf("starting server on port:%s", *addr)
    log.Fatal(http.ListenAndServe(*addr, mux))
}
