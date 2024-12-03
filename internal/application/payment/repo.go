package payment

import "trade-microservice.fyerfyer.net/internal/application/domain"

type Repository interface {
	GetCustomerByID(customerID uint64) (*domain.Customer, error)
	UpdateCustomerBalance(customer *domain.Customer) error
	SavePayment(payment *domain.Payment) error
}
