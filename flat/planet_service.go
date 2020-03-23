package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func Index() (*mongo.Cursor, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := ConnectDB().Find(ctx, bson.M{})
	return cursor, err
}

func Show(f bson.M) (Planet, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var planet Planet
	err := ConnectDB().FindOne(ctx, f).Decode(&planet)
	return planet, err
}

func Delete(f bson.M) (*mongo.DeleteResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := ConnectDB().DeleteOne(ctx, f)
	return result, err
}

func Store(p Planet) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := ConnectDB().InsertOne(ctx, p)
	return result, err
}
