package application

import (
	"trade-microservice.fyerfyer.net/internal/adapters/order"
	"trade-microservice.fyerfyer.net/internal/adapters/payment"
)

// we use Application to do dependency injection
type Application struct {
	order   *order.Adapter
	payment *payment.Adapter
}

func NewApplication(orderAdapter *order.Adapter, paymentAdapter *payment.Adapter) *Application {
	return &Application{
		order:   orderAdapter,
		payment: paymentAdapter,
	}
}
