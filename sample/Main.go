package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/translate/{level}", PostPhrase).Methods("POST")
	router.HandleFunc("/api/phrases/{level}", GetPhrase).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9999", router))
}