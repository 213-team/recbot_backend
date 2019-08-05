package main

import (
	"fmt"
	"io"
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
	// c.AddSubscription(
	// 	context.Background(),
	// 	&subscriptionb.AddSubscriptionReq{
	// 		Subscription: &subscriptionb.Subscription{
	// 			Channel: &subscriptionb.Channel{
	// 				Id: "123",
	// 			},
	// 			User: &subscriptionb.User{
	// 				Id: "321",
	// 			},
	// 		},
	// 	},
	// )
	// c.AddSubscription(
	// 	context.Background(),
	// 	&subscriptionb.AddSubscriptionReq{
	// 		Subscription: &subscriptionb.Subscription{
	// 			Channel: &subscriptionb.Channel{
	// 				Id: "1234",
	// 			},
	// 			User: &subscriptionb.User{
	// 				Id: "321",
	// 			},
	// 		},
	// 	},
	// )
	// c.AddSubscription(
	// 	context.Background(),
	// 	&subscriptionb.AddSubscriptionReq{
	// 		Subscription: &subscriptionb.Subscription{
	// 			Channel: &subscriptionb.Channel{
	// 				Id: "12345",
	// 			},
	// 			User: &subscriptionb.User{
	// 				Id: "3212",
	// 			},
	// 		},
	// 	},
	// )
	stream, err := c.ListSubscriptions(context.Background(), &subscriptionb.ListSubscriptionsReq{User: &subscriptionb.User{Id: "321"}})
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unknown internal error: %v", err)
		}
		fmt.Println(res.GetSubscription().User.Id, res.GetSubscription().Channel.Id)
	}
}
