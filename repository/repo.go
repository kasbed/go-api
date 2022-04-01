package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()


func connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil { log.Fatal(err) }
	err = client.Ping(ctx, nil)
	if err != nil { log.Fatal(err) }

	collection = client.Database("goApi").Collection("goApi")
}