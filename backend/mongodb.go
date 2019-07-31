package backend

import (
	"context"
	"fmt"
	"log"

	"github.com/213-team/recbot_backend/dao/mongodb"
	"github.com/213-team/recbot_backend/subscriptionb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test() {
	mongoCtx := context.Background()
	log.Println("Connecting...")
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Check whether the connection was succesful by pinging the MongoDB server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	// Bind our collection to our global variable for use in other methods
	subscription := subscriptionb.Subscription{
		Channel: &subscriptionb.Channel{Id: "CatsChannel1"},
		User:    &subscriptionb.User{Id: "FakeUser"}}
	mongodb.SubscriptionDAOMongoDB{}.Add(&subscription)

	subscription2 := subscriptionb.Subscription{
		Channel: &subscriptionb.Channel{Id: "CatsChannel2"},
		User:    &subscriptionb.User{Id: "FakeUser"}}
	mongodb.SubscriptionDAOMongoDB{}.Add(&subscription2)

	// data := models.Subscription{
	// 	ChannelID: "channel_id_1",
	// 	UserID:    "user_id_1",
	// }

	// _, err = blogdb.InsertOne(mongoCtx, data)
	// log.Println(err)

	// subscriptions, err2 := mongodb.SubscriptionDAOMongoDB{}.GetAllByUserID("FakeUser")

	// fetchedResult := blogdb.FindOne(mongoCtx, bson.M{"channel.id": "CatsChannel"})
	// if err2 != nil {
	// 	log.Fatalf("ERROR %v\n", err2)
	// }
	// for _, subscriptionCur := range subscriptions {
	// 	log.Println(subscriptionCur.Channel.Id, subscriptionCur.User.Id)
	// }
}
