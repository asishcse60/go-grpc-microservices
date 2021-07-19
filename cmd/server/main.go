package main

import (
	"fmt"
	"github.com/asishcse60/go-grpc-microservices/internal/db"
	"github.com/asishcse60/go-grpc-microservices/internal/rocket"
	"log"
)

func Run() error{
	rocketStore, err := db.NewDatabase()
	if err != nil{
		fmt.Println("Cannot connect db!")
		return err
	}
	err = rocketStore.Migrate()
	if err != nil{
		fmt.Println("database migrate failed!")
		return err
	}
	_=rocket.New(rocketStore)
	return nil
}
func main(){
	if err:=Run();err!=nil{
		log.Fatal(err)
	}
}