package payment

import (
	"errors"

	"trade-microservice.fyerfyer.net/internal/application/customer"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/ports"
)

type Service struct {
	repo  Repository
	cache ports.Cache
}

func NewService(repo Repository, cache ports.Cache) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
	}
}

// the totalAmount is passed by request order
func (s *Service) Charge(customerID uint64, orderID uint64, totalAmount float32) (*domain.Payment, error) {
	// first look up in cache
	var c *domain.Customer
	var success bool
	var err error
	success, c = customer.LookUpCustomerInCache(s.cache, customerID)
	if !success {
		c, err = s.repo.GetCustomerByID(customerID)
		if err != nil {
			return &domain.Payment{
				Status:  "failure",
				Message: "customer not found",
			}, err
		}
	}

	if c == nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "customer not found",
		}, errors.New("customer not found")
	}

	// purchase
	if !c.DeductBalance(totalAmount) {
		return &domain.Payment{
			Status:  "failure",
			Message: "customer didn't have enough money",
		}, errors.New("insufficient balance")
	}

	// to ensure data consistency, update customer cache manually
	if err := s.cache.Set(customer.GetCustomerKey(customerID), c, 0); err != nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "failed to update cache",
		}, err
	}

	// update customer database
	if err := s.repo.UpdateCustomerBalance(c); err != nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "failed to update database",
		}, err
	}

	payment := &domain.Payment{
		CustomerID: customerID,
		OrderID:    orderID,
		TotalPrice: totalAmount,
		Status:     "success",
		Message:    "successfully process payment",
	}

	if err := s.repo.SavePayment(payment); err != nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "failed to save completed payment",
		}, err
	}

	return payment, nil
}
