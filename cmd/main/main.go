package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database:%v", err)
	}

	// model migration
	err = db.AutoMigrate(&customerAdapter.Customer{})
	if err != nil {
		log.Fatalf("failed to migrate customers: %v", err)
	}

	err = db.AutoMigrate(&orderAdapter.Order{}, &orderAdapter.OrderItem{})
	if err != nil {
		log.Fatalf("failed to migrate orders: %v", err)
	}

	err = db.AutoMigrate(&paymentAdapter.Payment{})
	if err != nil {
		log.Fatalf("failed to migrate payments: %v", err)
	}

	// setup repo
	orderRepo := orderAdapter.NewGormRepository(db)
	paymentRepo := paymentAdapter.NewGormRepository(db)
	customerRepo := customerAdapter.NewGormRepository(db)

	// setup cache
	redisClient := redis.NewRedisClient("127.0.0.1:6379", "", 10, 10, 3*time.Minute)
	log.Println("successfully set up redis connection")

	// setup service
	paymentService := paymentService.NewService(paymentRepo, redisClient)
	orderService := orderService.NewService(orderRepo, paymentService)
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
