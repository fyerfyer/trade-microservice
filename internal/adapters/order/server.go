package order

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"trade-microservice.fyerfyer.net/internal/application/order"
	pb "trade-microservice.fyerfyer.net/proto/order"
)

type Adapter struct {
	service *order.Service
	port    int
	pb.UnimplementedOrderServer
}

func NewAdapter(service *order.Service, port int) *Adapter {
	return &Adapter{
		service: service,
		port:    port,
	}
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d:%v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServer(grpcServer, a)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d:%v", a.port, err)
	}
}
