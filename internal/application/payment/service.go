package payment

import (
	"errors"
	"time"

	"trade-microservice.fyerfyer.net/internal/application/domain"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// the totalAmount is passed by request order
func (s *Service) Charge(customerID uint64, orderID uint64, totalAmount float32) (*domain.Payment, error) {
	customer, err := s.repo.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}

	// purchase
	if customer.Balance < totalAmount {
		return &domain.Payment{
			Status:  "failure",
			Message: "customer didn't have enough money",
		}, errors.New("insufficient balance")
	}

	customer.DeductBalance(totalAmount)
	if err := s.repo.UpdateCustomerBalance(customer); err != nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "failed to update customer status",
		}, err
	}

	payment := &domain.Payment{
		CustomerID: customerID,
		OrderID:    orderID,
		TotalPrice: totalAmount,
		Status:     "success",
		CreatedAt:  time.Now(),
	}

	if err := s.repo.SavePayment(payment); err != nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "failed to save completed payment",
		}, err
	}

	return payment, nil
}
