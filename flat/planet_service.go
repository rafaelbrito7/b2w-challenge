package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Index() (*mongo.Cursor, error) {
	cursor, err := ConnectDB().Find(context.TODO(), bson.M{})
	return cursor, err
}

func Show(f bson.M) (Planet, error) {
	var planet Planet
	err := ConnectDB().FindOne(context.TODO(), f).Decode(&planet)
	return planet, err
}

func Delete(f bson.M) (*mongo.DeleteResult, error) {
	result, err := ConnectDB().DeleteOne(context.TODO(), f)
	return result, err
}
