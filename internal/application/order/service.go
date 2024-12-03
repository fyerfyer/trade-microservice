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

func (s *Service) CreateOrder(customerID uint64, items []domain.OrderItem) (*domain.Order, error) {
	order := &domain.Order{
		CustomerID: customerID,
		Items:      items,
		Status:     "unprocessed",
	}

	err := s.repo.Save(order)
	if err != nil {
		return nil, err
	}

	// pay for the order
	totalPrice := order.TotalPrice()
	res, err := s.payment.Charge(customerID, order.ID, totalPrice)

	// case 1: payment error
	if err != nil {
		order.Status = "unpaid"
		s.repo.Update(order)
		return nil, errors.New("failed to process payment")
	}
	// case 2: successfully paid
	if res.Status == "success" {
		s.repo.Delete(order.ID)
		return order, nil
	}

	order.Status = "unpaid"
	s.repo.Update(order)
	return nil, errors.New(res.Message)
}

func (s *Service) GetUnpaidOrders(userID uint64) ([]domain.Order, error) {
	return s.repo.FindUnpaidByUser(userID)
}
