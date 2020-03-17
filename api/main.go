package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting the application...")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	router := mux.NewRouter()

	router.HandleFunc("/api/planets", GetPlanets).Methods("GET")
	router.HandleFunc("/api/planets/{id}", GetPlanet).Methods("GET")
	//TODO router.HandleFunc("/api/planets", createPlanet).Methods("POST")
	router.HandleFunc("/api/planets/{id}", UpdatePlanet).Methods("PUT")
	//TODO router.HandleFunc("/api/planets/{id}", deletePlanet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
