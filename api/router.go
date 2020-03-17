package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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

func GetPlanets(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var planets []Planet
	planets, err := Index()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, planets)
}

func GetPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	var planet Planet
	planet, err := Show(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	respondWithJson(w, http.StatusOK, planet)
}

func UpdatePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	defer r.Body.Close()
	params := mux.Vars(r)
	var planet Planet

	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Body")
		return
	}
	planet, err := Update(params["id"], planet)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, planet)
}

func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	defer r.Body.Close()
	var params = mux.Vars(r)

	result, err := Delete(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, result)
}

//func CreatePlanet(response http.ResponseWriter, request *http.Request) {
//	response.Header().Add("content-type", "application/json")
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//
//	var planet models.Planet
//	var SWAPIdata models.SWAPIresponse
//
//	_ = json.NewDecoder(request.Body).Decode(&planet)
//
//	values := url.Values{
//		"search": []string {planet.Name},
//	}
//
//	rs, err := http.Get("https://swapi.co/api/planets/?" + values.Encode())
//	if err != nil {
//		fmt.Printf("The HTTP request failed with error %s\n", err)
//	}
//	defer rs.Body.Close()
//	data, _ := ioutil.ReadAll(rs.Body)
//	_ = json.Unmarshal(data, &SWAPIdata)
//
//	if SWAPIdata.Count == 0 {
//		json.NewEncoder(response).Encode("Planet does not exists")
//		return
//	}
//
//	planet.FilmsApparitionsCount = len(SWAPIdata.Results[0].Films)
//
//	collection := ConnectDB()
//
//	result, err  := collection.InsertOne(ctx, planet)
//	if err != nil {
//		fmt.Printf("error")
//	}
//
//	json.NewEncoder(response).Encode(result)
//}