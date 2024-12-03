package order

import (
	"errors"

	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/ports"
)

type Service struct {
	repo    Repository
	payment ports.PaymentPort
}

func NewService(repo Repository, paymentPort ports.PaymentPort) *Service {
	return &Service{
		repo:    repo,
		payment: paymentPort,
	}
}

func (s *Service) ProcessOrder(customerID uint64, items []domain.OrderItem) error {
	order := &domain.Order{
		CustomerID: customerID,
		Items:      items,
		Status:     "unprocessed",
	}

	err := s.repo.Save(order)
	if err != nil {
		return err
	}

	// pay for the order
	totalPrice := order.TotalPrice()
	res, err := s.payment.Charge(customerID, order.ID, totalPrice)

	// successfully paid
	if err == nil && res.Status == "success" {
		s.repo.Delete(order.ID)
		return nil
	}

	order.Status = "unpaid"
	s.repo.Update(order)
	return errors.New(res.Message)
}
