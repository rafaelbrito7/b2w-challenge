package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting the application...")

	router := mux.NewRouter()

	router.HandleFunc("/api/totalPlanets", GetTotalPlanets).Methods("GET")
	router.HandleFunc("/api/planets/{id}", GetPlanetById).Methods("GET")
	router.HandleFunc( "/api/planets", GetPlanetByName).Methods("GET")
	router.HandleFunc("/api/planets", CreatePlanet).Methods("POST")
	router.HandleFunc("/api/planets/{id}", DeletePlanet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Finishing the application...")

}
