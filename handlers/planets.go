package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var url = "https://swapi.co/planets"

type Planet struct {
	Name      string   `json:"name"`
	URL       string   `json:"url"`
	Residents []string `json:"residents"`
	Films     []string `json:"films"`
}

var client = &http.Client{Timeout: 10 * time.Second}

func NewPlanets() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/planets").Subrouter()

	router.HandleFunc("/", List)

	return router
}

func List(res http.ResponseWriter, req *http.Request) {
	r, err := client.Get(url)

	if err != nil {
		log.Panic("Error to list planets", err)
	}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	var planet Planet

	if err := decoder.Decode(&planet); err != nil {
		log.Panic("Fail to convert Panets", err)
	}

	println("Planets are %s", &planet)
}
