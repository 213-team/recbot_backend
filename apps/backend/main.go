package main

import (
	"fmt"
	"log"
	"net"

	"github.com/213-team/recbot_backend/backend"
	"github.com/213-team/recbot_backend/subscriptionb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func loadConfig() {
	viper.SetConfigName("default")
	viper.AddConfigPath("$HOME/go/src/github.com/213-team/recbot_backend")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	loadConfig()
	backendPort := viper.GetString("backend.port")
	fmt.Printf("Starting server on port :%s...", backendPort)
	listener, err := net.Listen("tcp", ":"+backendPort)
	if err != nil {
		log.Fatalf("Unable to listen on port :%s: %v", backendPort, err)
	}

	opts := []grpc.ServerOption{}

	server := &backend.SubscriptionService{}
	grpcServer := grpc.NewServer(opts...)

	subscriptionb.RegisterSubscriptionServiceServer(grpcServer, server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
