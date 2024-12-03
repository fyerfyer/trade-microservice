package ports

import "trade-microservice.fyerfyer.net/internal/application/domain"

type OrderPort interface {
	ProcessOrder(customerID uint64, items []domain.OrderItem) error
}
