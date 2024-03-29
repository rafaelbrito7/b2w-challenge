package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (p *Planet) validate() url.Values{
	errs := url.Values{}

	if p.Name == "" {
		errs.Add("requestBody", "Error: Missing planet's name")
	}
	if p.Climate == "" {
		errs.Add("requestBody", "Error: Missing planet's climate")
	}
	if p.Terrain == "" {
		errs.Add("requestBody", "Error: Missing planet's terrain")
	}
	return errs
}