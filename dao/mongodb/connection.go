package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetConnection gets a connection to a mongodb
func GetConnection() (*mongo.Client, context.Context) {
	mongoCtx := context.Background()
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	}

	return db, mongoCtx
}

//GetSubscriptionCollection gets a handle for "subscriptions" collection
func GetSubscriptionCollection() (*mongo.Collection, context.Context) {
	db, mongoCtx := GetConnection()
	return db.Database("recbot").Collection("subscriptions"), mongoCtx
}
