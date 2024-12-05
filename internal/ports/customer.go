package ports

import "trade-microservice.fyerfyer.net/internal/application/domain"

type CustomerPort interface {
	CreateCustomer(name string) (*domain.Customer, error)
	GetCustomer(id uint64) (*domain.Customer, error)
	DeactiveCustomer(id uint64) error
	ReactiveCustomer(id uint64) error
	SubmitOrder(customerID uint64, items []domain.OrderItem) (bool, error)
	GetUnpaidOrders(customerID uint64) ([]domain.Order, error)
	StoreBalance(customerID uint64, balance float32) error
}
