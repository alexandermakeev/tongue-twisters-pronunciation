package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", http.FileServer(http.Dir("static")))
	router.HandleFunc("/api/translate", PostPhrase)
	router.HandleFunc("/api/phrases/{level}", GetPhrase)

	log.Fatal(http.ListenAndServe(":9999", router))
}