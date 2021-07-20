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
	return &rkt.GetRocketResponse{}, nil
}

func (h Handler) AddRocket(ctx context.Context, r *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	return &rkt.AddRocketResponse{}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, r *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	return &rkt.DeleteRocketResponse{}, nil
}

func New(rktService RocketService) Handler  {
	return Handler{RocketService: rktService}
}