package repository

import (
	"context"
	"log"
	"src/structs"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()


func connect() * mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil { log.Fatal(err) }
	err = client.Ping(ctx, nil)
	if err != nil { log.Fatal(err) }

	collection = client.Database("goApi").Collection("goApi")

	return client
}

func insertUser(user *structs.UserDTO) *structs.UserDTO {
	//Convert dto to data and insert in mongo
	client := connect();
	collection = client.Database("goApi").Collection("user")
	data := structs.User{ID: primitive.NewObjectID(), CreatedAt: time.Now(), UpdatedAt: time.Now(), 
		UserName: user.UserName, Email: user.Email, FirstName: user.FirstName, SurName: user.SurName}
	collection.InsertOne(ctx, data)
	return user;
}

func removeUser(id *primitive.ObjectID) {
	client := connect();
	collection = client.Database("goApi").Collection("user")
	collection.DeleteOne(ctx, id)
}

