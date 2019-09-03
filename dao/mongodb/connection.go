package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s", viper.GetString("mongo.host"), viper.GetString("mongo.port"))
}

//GetConnection gets a connection to a mongodb
func GetConnection() (*mongo.Client, context.Context) {
	mongoCtx := context.Background()
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(getMongoConnectionString()))
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
	dbname := viper.GetString("mongo.dbname")
	return db.Database(dbname).Collection("subscriptions"), mongoCtx
}
