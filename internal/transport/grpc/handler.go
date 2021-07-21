package grpc

import (
	"context"
	rkt "github.com/asishcse60/go-grpc-microservices/grpc-protos/rocket/v1"
	"github.com/asishcse60/go-grpc-microservices/internal/rocket"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RocketService interface {
	GetRocketByID(ctx context.Context,  id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}
type Handler struct {
	RocketService RocketService
}

func (h Handler) Serve() error{
	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Print(err)
		return err
	}
	grpcServer := grpc.NewServer()
    rkt.RegisterRocketServiceServer(grpcServer, &h)
	if err := grpcServer.Serve(lis); err!=nil{
		log.Println(err)
		return err
	}
	return nil
}

func (h Handler) GetRocket(ctx context.Context, r *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Println("Get rocket grpc endpoint hit")
	rocket, err := h.RocketService.GetRocketByID(ctx, r.Id)
	if err !=nil{
		log.Println(err)
		return &rkt.GetRocketResponse{}, err}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id: rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, r *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {

	log.Println("Add rocket grpc endpoint hits")

	rckt ,err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID: r.Rocket.Id,
		Name: r.Rocket.Name,
		Type: r.Rocket.Type,
	})
	if err != nil{
		log.Println(err.Error())
		return nil, err
	}
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id: rckt.ID,
			Name: rckt.Name,
			Type: rckt.Type,
		},
	}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, r *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Println("Delete rocket grpc endpoint hits")
	err := h.RocketService.DeleteRocket(ctx, r.Rocket.Id)
	if err!=nil{
		log.Println(err.Error())
		return nil, err
	}
	return &rkt.DeleteRocketResponse{Status: "Successfully delete rocket"},nil
}

func New(rktService RocketService) Handler  {
	return Handler{RocketService: rktService}
}