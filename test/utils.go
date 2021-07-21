package test

import (
	rkt "github.com/asishcse60/go-grpc-microservices/grpc-protos/rocket/v1"
	"google.golang.org/grpc"
	"log"
)

func GetClient() rkt.RocketServiceClient{
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err!=nil{
		log.Fatalf("could not connect: %s", err)
	}
	rocketClient := rkt.NewRocketServiceClient(conn)
	return rocketClient
}
