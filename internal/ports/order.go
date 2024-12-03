package ports

import "trade-microservice.fyerfyer.net/internal/application/domain"

type OrderPort interface {
	CreateOrder(customerID uint64, items []domain.OrderItem) (*domain.Order, error)
	GetUnpaidOrders(customerID uint64) ([]domain.Order, error)
}
