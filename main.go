package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/mpaivafontes/star-wars-go/handlers"
)

func main() {
	router := mux.NewRouter()

	planet := planets.NewPlanets()

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
