package main

import (
	"log"
	"time"

	customerAdapter "trade-microservice.fyerfyer.net/internal/adapters/customer"
	orderAdapter "trade-microservice.fyerfyer.net/internal/adapters/order"
	paymentAdapter "trade-microservice.fyerfyer.net/internal/adapters/payment"
	redis "trade-microservice.fyerfyer.net/internal/application/cache"
	customerService "trade-microservice.fyerfyer.net/internal/application/customer"
	orderService "trade-microservice.fyerfyer.net/internal/application/order"
	paymentService "trade-microservice.fyerfyer.net/internal/application/payment"
)

func main() {
	dsn := "root:110119abc@tcp(127.0.0.1:3306)/microservice?charset=utf8&parseTime=true"

	// setup repo
	var (
		customerRepo *customerAdapter.GormRepository
		orderRepo    *orderAdapter.GormRepository
		paymentRepo  *paymentAdapter.GormRepository
		err          error
	)

	if customerRepo, err = customerAdapter.NewGormRepository(dsn); err != nil {
		log.Fatalf("failed to init customer database: %v", err)
	}

	if orderRepo, err = orderAdapter.NewGormRepository(dsn); err != nil {
		log.Fatalf("failed to init order database: %v", err)
	}

	if paymentRepo, err = paymentAdapter.NewGormRepository(dsn); err != nil {
		log.Fatalf("failed to init payment database: %v", err)
	}

	// setup cache
	redisClient := redis.NewRedisClient("127.0.0.1:6379", "", 10, 10, 3*time.Minute)
	log.Println("successfully set up redis connection")

	// setup service
	paymentService := paymentService.NewService(paymentRepo, redisClient)
	orderService := orderService.NewService(orderRepo, redisClient, paymentService)
	customerService := customerService.NewService(customerRepo, redisClient, orderService)

	// setup adapter
	paymentAdapter := paymentAdapter.NewAdapter(paymentService, 50051)
	orderAdapter := orderAdapter.NewAdapter(orderService, 50052)

	// start backend service
	go func() {
		paymentAdapter.Run()
	}()

	go func() {
		orderAdapter.Run()
	}()

	// expose customer service for user to invoke
	customerAdapter := customerAdapter.NewAdapter(customerService, 50053)
	log.Println("customer server running on port 50053...")
	customerAdapter.Run()
}
