package customer

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"trade-microservice.fyerfyer.net/internal/application/customer"
	"trade-microservice.fyerfyer.net/internal/tls"
	pb "trade-microservice.fyerfyer.net/proto/customer"
)

type Adapter struct {
	service *customer.Service
	port    int
	pb.UnimplementedCustomerServer
}

func NewAdapter(service *customer.Service, port int) *Adapter {
	return &Adapter{
		service: service,
		port:    port,
	}
}

func (a *Adapter) Run() {
	tlsCredentials, err := tls.GetServerTLSCredentials()
	if err != nil {
		log.Fatalf("tls config failed: %v", err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d:%v", a.port, err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(tlsCredentials))

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCustomerServer(grpcServer, a)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d:%v", a.port, err)
	}
}
