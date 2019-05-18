package main

import (
	"log"

	"github.com/213-team/recbot_backend/subscriptionb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := subscriptionb.NewSubscriptionServiceClient(conn)
	response, err := c.ReadSubscription(context.Background(), &subscriptionb.ReadSubscriptionReq{User: &subscriptionb.User{Id: "333"}})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Subscription.User.GetId())
}
