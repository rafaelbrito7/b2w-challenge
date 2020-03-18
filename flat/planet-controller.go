package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

//TODO create GetPlanetsByName POST method
//TODO check every single method's validation process
//TODO finish as soon as possible so I can start unit tests

func GetSwapiPlanet(name string) (*http.Response, error)  {
	values := url.Values{
		"search": []string {name},
	}
	rs, err := http.Get("https://swapi.co/api/planets/?" + values.Encode())

	return rs, err
}

func GetPlanets(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var planets []Planet
	cursor, err := Index()
	for cursor.Next(ctx) {
		var planet Planet
		err := cursor.Decode(&planet)
		if err != nil {
			log.Fatal(err)
		}
		planets = append(planets, planet)
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(planets) == 0 {
		respondWithError(w, http.StatusBadRequest, "No planets registered!")
		return
	}
	respondWithJson(w, http.StatusOK, planets)
}

func GetPlanetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	var planet Planet
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	planet, err := Show(filter)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Planet not found!")
		return
	}
	respondWithJson(w, http.StatusOK, planet)
}

func GetPlanetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	var planet Planet
	planet.Name = params["name"]
	filter := bson.M{"name": planet.Name}
	result := ConnectDB().FindOne(context.TODO(), filter)
	respondWithJson(w, http.StatusOK, result)
}

func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	defer r.Body.Close()
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	result, err := Delete(filter)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if result.DeletedCount == 0 {
		respondWithError(w, http.StatusBadRequest, "Planet not found!")
		return
	}
	respondWithJson(w, http.StatusOK, "Planet deleted with success!")
}

func CreatePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var planet Planet
	var swapiresponse SWAPIresponse

	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	if validErrs := planet.validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	rs, err := GetSwapiPlanet(planet.Name)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	defer rs.Body.Close()
	data, _ := ioutil.ReadAll(rs.Body)
	_ = json.Unmarshal(data, &swapiresponse)

	if swapiresponse.Count == 0 {
		respondWithJson(w, http.StatusBadRequest, "Planet does not exist")
		return
	}

	planet.FilmsApparitionsCount = len(swapiresponse.Results[0].Films)
	result, err  := ConnectDB().InsertOne(ctx, planet)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJson(w, http.StatusOK, result)
}