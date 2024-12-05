package payment

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"trade-microservice.fyerfyer.net/internal/application/payment"
	pb "trade-microservice.fyerfyer.net/proto/payment"
)

type Adapter struct {
	service *payment.Service
	port    int
	pb.UnimplementedPaymentServer
}

func NewAdapter(service *payment.Service, port int) *Adapter {
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
	pb.RegisterPaymentServer(grpcServer, a)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d:%v", a.port, err)
	}
}
