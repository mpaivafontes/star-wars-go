package main

import (
	"log"
	"net/http"
)

type Planet struct {
	Name      string   `json:"name"`
	URL       string   `json:"url"`
	Residents []string `json:"residents"`
	Films     []string `json:"films"`
}

func FindAllPlanets(w http.ResponseWriter, r *http.Request) String {
	url := Environment.BaseURL + "/planets"

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("fail to find planets", err)
	}

	return resp.Body.Read
}
