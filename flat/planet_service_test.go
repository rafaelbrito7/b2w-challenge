package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/totalPlanets", GetTotalPlanets).Methods("GET")
	return router
}

func TestGetTotalPlanets(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/totalPlanets", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "StatusOK and all documents are expected.")
}