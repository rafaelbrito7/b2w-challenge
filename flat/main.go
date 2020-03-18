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

	router.HandleFunc("/api/planets", GetPlanets).Methods("GET")
	router.HandleFunc("/api/planets/{id}", GetPlanetById).Methods("GET")
	router.HandleFunc( "api/planets/{name}", GetPlanetByName).Methods("GET")
	router.HandleFunc("/api/planets", CreatePlanet).Methods("POST")
	// router.HandleFunc("/flat/planets/{id}", UpdatePlanet).Methods("PUT")
	router.HandleFunc("/api/planets/{id}", DeletePlanet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Finishing the application...")

}
