package main

import (
	"log"

	"github.com/time_capsule/Api-Gateway-Timecapsule/api" // Your API package
	"github.com/time_capsule/Api-Gateway-Timecapsule/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()

	// gRPC connections
	timelineGrpcConn, err := grpc.NewClient(cfg.TimelineSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Timeline service: %v", err)
	}
	defer timelineGrpcConn.Close()

	// memoryGrpcConn, err := grpc.NewClient(cfg.MemorySvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Timeline service: %v", err)
	// }
	// defer memoryGrpcConn.Close()

	memoryGrpcConn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Timeline service: %v", err)
	}
	defer memoryGrpcConn.Close()

	// Create router
	router := api.NewRouter(timelineGrpcConn, memoryGrpcConn)

	// Start server
	if err := router.Run(cfg.HTTPPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
