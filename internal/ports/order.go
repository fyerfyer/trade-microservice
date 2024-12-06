package ports

import "trade-microservice.fyerfyer.net/internal/application/domain"

type OrderPort interface {
	ProcessItems(customerID uint64, items []domain.OrderItem) error
	ProcessOrder(order *domain.Order) error
}
