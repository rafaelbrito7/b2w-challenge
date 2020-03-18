package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	)

type Planet struct {
ID  										primitive.ObjectID  			`json:"_id,omitempty" bson:"_id,omitempty"`
Name  										string  						`json:"name,omitempty" bson:"name,omitempty"`
Climate  									string  						`json:"climate,omitempty" bson:"climate,omitempty"`
Terrain										string 							`json:"terrain,omitempty" bson:"terrain,omitempty"`
FilmsApparitionsCount 						int								`json:"filmsapparitionscount,omitempty" bson:"filmsapparitionscount,omitempty"`
}
  
type SWAPIresponse struct {
        Count    int         `json:"count"`
        Results  []struct {
                Name           string    `json:"name"`
                Climate        string    `json:"climate"`
                Terrain        string    `json:"terrain"`
                Films          []string  `json:"films"`
        } `json:"results"`
}
