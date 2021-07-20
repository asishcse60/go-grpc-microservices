package main

import (
	"fmt"
	"github.com/asishcse60/go-grpc-microservices/internal/db"
	"github.com/asishcse60/go-grpc-microservices/internal/rocket"
	"github.com/asishcse60/go-grpc-microservices/internal/transport/grpc"
	"log"
)

func Run() error {
	rocketStore, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Cannot connect db!")
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		fmt.Println("database migrate failed!")
		return err
	}
	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(&rktService)
	if err != nil{
		return err
	}
    if err := rktHandler.Serve(); err != nil {
		return err
	}
	return nil
}
func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
