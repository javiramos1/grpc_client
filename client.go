package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/javiramos1/grpcapi"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Starting client...")

	hostname := os.Getenv("SVC_HOST_NAME")

	if len(hostname) <= 0 {
		hostname = "0.0.0.0"
	}

	port := os.Getenv("SVC_PORT")

	if len(port) <= 0 {
		port = "50051"
	}

	cc, err := grpc.Dial(hostname+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := grpcapi.NewGrpcServiceClient(cc)
	fmt.Printf("Created client: %f", c)

	callService(c)

}

func callService(c grpcapi.GrpcServiceClient) {
	fmt.Println("callService...")
	req := &grpcapi.GrpcRequest{
		Input: "test",
	}
	res, err := c.GrpcService(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling gRPC: %v", err)
	}
	log.Printf("Response from Service: %v", res.Response)
}
