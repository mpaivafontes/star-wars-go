package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var(
	URL := E
)

func main() {
	router := mux.NewRouter()

	


	router.HandleFunc("/", P.Find)

	log.Fatal(http.ListenAndServe(":8000", router))
}
