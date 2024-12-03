package ports

import "trade-microservice.fyerfyer.net/internal/application/domain"

type PaymentPort interface {
	Charge(customerID uint64, orderID uint64, totalAmount float32) (*domain.Payment, error)
}
