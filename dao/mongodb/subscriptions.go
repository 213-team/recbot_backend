package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/213-team/recbot_backend/dao/interfaces"
	"github.com/213-team/recbot_backend/subscriptionb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//SubscriptionDAOMongoDB is a mongodb impl for SubscriptionDAO
type SubscriptionDAOMongoDB struct {
}

//Add adds subscription to a database CHANGE
func (dao SubscriptionDAOMongoDB) Add(subscription *subscriptionb.Subscription) error {
	subscriptionsColl, mongoCtx := GetSubscriptionCollection()
	_, err := subscriptionsColl.InsertOne(mongoCtx, subscription)
	return err
}

//Delete removes a subscription from a database
func (dao SubscriptionDAOMongoDB) Delete(subscription *subscriptionb.Subscription) error {
	subscriptionsColl, mongoCtx := GetSubscriptionCollection()
	filter := bson.M{"user.id": subscription.User.Id, "channel.id": subscription.Channel.Id}
	_, err := subscriptionsColl.DeleteOne(mongoCtx, filter)
	return err
}

//GetOne gers subscrtiption from the database
func (dao SubscriptionDAOMongoDB) GetOne(userID string, channelID string) (*subscriptionb.Subscription, error) {
	subscriptionsColl, mongoCtx := GetSubscriptionCollection()
	filter := bson.M{"user.id": userID, "channel.id": channelID}
	foundSubscription := subscriptionb.Subscription{}
	err := subscriptionsColl.FindOne(mongoCtx, filter).Decode(&foundSubscription)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find subscription userID=%s channelID=%s %v", userID, channelID, err))
	}
	return &foundSubscription, nil
}

//GetAllByUserID gets all subscriptions by userID
func (dao SubscriptionDAOMongoDB) GetAllByUserID(userID string) (interfaces.SubscriptionCursor, error) {
	subscriptionsColl, mongoCtx := GetSubscriptionCollection()
	filter := bson.M{"user.id": userID}

	cur, err := subscriptionsColl.Find(mongoCtx, filter)
	if err != nil {
		log.Fatal(err)
	}
	subscriptionCursor := SubscriptionCursorMongoDB{
		cur,
		mongoCtx,
	}
	return subscriptionCursor, nil
}

//SubscriptionCursorMongoDB is an implementation of SubscriptionCursor iface
type SubscriptionCursorMongoDB struct {
	mongoCursor  *mongo.Cursor
	mongoContext context.Context
}

//Next is an implementation of SubscriptionCursor iface
func (cursor SubscriptionCursorMongoDB) Next() (*subscriptionb.Subscription, error) {
	if !cursor.mongoCursor.Next(cursor.mongoContext) {
		defer cursor.mongoCursor.Close(cursor.mongoContext)
		if err := cursor.mongoCursor.Err(); err != nil {
			log.Fatal(err)
		}
		return nil, iterator.Done
	}
	var result subscriptionb.Subscription
	err := cursor.mongoCursor.Decode(&result)
	if err != nil {
		defer cursor.mongoCursor.Close(cursor.mongoContext)
		log.Fatal(err)
	}
	return &result, nil
}
