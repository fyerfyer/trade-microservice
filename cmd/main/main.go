package main

import (
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trade-microservice.fyerfyer.net/internal/adapters/customer"
	customerModel "trade-microservice.fyerfyer.net/internal/adapters/customer"
	customerRepo "trade-microservice.fyerfyer.net/internal/adapters/customer"
	"trade-microservice.fyerfyer.net/internal/adapters/order"
	orderModel "trade-microservice.fyerfyer.net/internal/adapters/order"
	orderRepo "trade-microservice.fyerfyer.net/internal/adapters/order"
	"trade-microservice.fyerfyer.net/internal/adapters/payment"
	paymentModel "trade-microservice.fyerfyer.net/internal/adapters/payment"
	paymentRepo "trade-microservice.fyerfyer.net/internal/adapters/payment"
	customerService "trade-microservice.fyerfyer.net/internal/application/customer"
	orderService "trade-microservice.fyerfyer.net/internal/application/order"
	paymentService "trade-microservice.fyerfyer.net/internal/application/payment"
	pb "trade-microservice.fyerfyer.net/proto/customer"
)

func main() {
	dsn := "root:110119abc@tcp(127.0.0.1:3306)/microservice?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database:%v", err)
	}

	// model migration
	err = db.AutoMigrate(&orderModel.Order{}, &customerModel.Customer{}, &paymentModel.Payment{})
	if err != nil {
		log.Fatalf("failed to migrate model:%v", err)
	}

	// setup repo
	orderRepo := orderRepo.NewGormRepository(db)
	paymentRepo := paymentRepo.NewGormRepository(db)
	customerRepo := customerRepo.NewGormRepository(db)

	// setup service
	paymentService := paymentService.NewService(paymentRepo)
	orderService := orderService.NewService(orderRepo, paymentService)
	customerService := customerService.NewService(customerRepo, orderService)

	// setup adapter
	paymentAdapter := payment.NewAdapter(paymentService, 50051)
	orderAdapter := order.NewAdapter(orderService, 50052)

	// start backend service
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		paymentAdapter.Run()
	}()

	go func() {
		defer wg.Done()
		orderAdapter.Run()
	}()

	wg.Wait()

	// expose customer service for user to invoke
	customerAdapter := customer.NewAdapter(customerService, 50053)
	listen, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on port 50053: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCustomerServer(grpcServer, customerAdapter)
	reflection.Register(grpcServer)

	log.Println("customer grpc server running on port 50053")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc server:%v", err)
	}
}
