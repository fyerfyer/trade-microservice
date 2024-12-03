package payment

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"trade-microservice.fyerfyer.net/internal/ports"
	pb "trade-microservice.fyerfyer.net/proto/payment"
)

type Adapter struct {
	api    ports.PaymentPort
	port   int
	server *grpc.Server
	pb.UnimplementedPaymentServer
}

func NewAdapter(api ports.PaymentPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d:%v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServer(grpcServer, a)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d:%v", a.port, err)
	}
}
