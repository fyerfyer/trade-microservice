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
	order := domain.NewOrder(customerID, items)

	err := s.repo.Save(order)
	if err != nil {
		return err
	}

	// pay for the order
	totalPrice := order.TotalPrice()
	res, err := s.payment.Charge(customerID, order.ID, totalPrice)

	// successfully paid
	if err == nil && res.Status == "success" {
		order.Status = "success"
		s.repo.Update(order)
		return nil
	}

	order.Status = "unpaid"
	s.repo.Update(order)
	return errors.New(res.Message)
}
