package customer

import (
	"errors"

	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/ports"
)

type Service struct {
	repo  Repository
	order ports.OrderPort
}

func NewService(repo Repository, orderPort ports.OrderPort) *Service {
	return &Service{
		repo:  repo,
		order: orderPort,
	}
}

func (s *Service) CreateCustomer(name string) (*domain.Customer, error) {
	customer := &domain.Customer{
		Name:    name,
		Status:  "active",
		Balance: 0,
	}

	err := s.repo.Save(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) GetCustomer(id uint64) (*domain.Customer, error) {
	return s.repo.GetByID(id)
}

func (s *Service) DeactiveCustomer(id uint64) error {
	customer, err := s.GetCustomer(id)
	if err != nil {
		return err
	}

	if customer.Status != "inactive" {
		customer.Status = "inactive"
		return s.repo.Update(customer)
	}

	return nil
}

func (s *Service) ReactiveCustomer(id uint64) error {
	customer, err := s.GetCustomer(id)
	if err != nil {
		return err
	}

	if customer.Status != "active" {
		customer.Status = "active"
		return s.repo.Update(customer)
	}

	return nil
}

func (s *Service) SubmitOrder(customerID uint64, items []domain.OrderItem) (bool, error) {
	customer, err := s.repo.GetByID(customerID)
	if err != nil {
		return false, errors.New("customer not found")
	}
	if customer.Status != "active" {
		return false, errors.New("customer is inactive")
	}

	// the CreateOrderMethod will process the order
	err = s.order.ProcessOrder(customerID, items)

	// failed to pay
	if err != nil {
		return false, err
	}

	return true, nil
}
