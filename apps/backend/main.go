package main

import (
	"fmt"
	"log"
	"net"

	"github.com/213-team/recbot_backend/backend"
	"github.com/213-team/recbot_backend/subscriptionb"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50051...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	opts := []grpc.ServerOption{}

	server := &backend.SubscriptionService{}
	grpcServer := grpc.NewServer(opts...)

	subscriptionb.RegisterSubscriptionServiceServer(grpcServer, server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
