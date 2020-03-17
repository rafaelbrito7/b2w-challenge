package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"

)

func Index() ([]Planet, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var planets []Planet
	cursor, err := ConnectDB().Find(context.TODO(), bson.M{})

	for cursor.Next(ctx) {
		var planet Planet
		err := cursor.Decode(&planet)
		if err != nil {
			log.Fatal(err)
		}
		planets = append(planets, planet)
	}
	return planets, err
}

func Show(i string) (Planet, error) {
	var planet Planet
	id, _ := primitive.ObjectIDFromHex(i)
	filter := bson.M{"_id": id}
	err := ConnectDB().FindOne(context.TODO(), filter).Decode(&planet)

	return planet, err
}

func Update(i string, p Planet) (Planet, error) {
	var planet Planet
	id, _ := primitive.ObjectIDFromHex(i)
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"name", p.Name},
			{"climate", p.Climate},
			{"terrain", p.Terrain},
			{"filmsapparitionscount", p.FilmsApparitionsCount},
		}},
	}
	err := ConnectDB().FindOneAndUpdate(context.TODO(), filter, update).Decode(&planet)
	planet.ID = id

	return planet, err
}

func Delete(i string) (*mongo.DeleteResult, error) {
	id, _ := primitive.ObjectIDFromHex(i)
	filter := bson.M{"_id": id}

	result, err := ConnectDB().DeleteOne(context.TODO(), filter)

	return result, err
}
