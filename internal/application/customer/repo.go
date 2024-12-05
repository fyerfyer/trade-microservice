package customer

import (
	"trade-microservice.fyerfyer.net/internal/application/domain"
)

type Repository interface {
	Save(customer *domain.Customer) error
	Update(customer *domain.Customer) error
	GetByID(customerID uint64) (*domain.Customer, error)
	GetByName(customerName string) (*domain.Customer, error)
	GetUnpaidOrdersByID(customerID uint64) ([]domain.Order, error)
}
