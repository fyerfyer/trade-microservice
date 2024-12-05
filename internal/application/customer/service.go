package customer

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/ports"
)

type Service struct {
	repo  Repository
	cache ports.Cache
	order ports.OrderPort
}

func NewService(repo Repository, cache ports.Cache, orderPort ports.OrderPort) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
		order: orderPort,
	}
}

func (s *Service) CreateCustomer(name string) (*domain.Customer, error) {
	// find in the cache first
	if s.cache.Exists(name) {
		return nil, errors.New("duplicate customer")
	}

	// else, find in db
	// if we can find in db, then update our cache
	// we store id and name key separately
	if c, _ := s.repo.GetByName(name); c != nil {
		cKey := GetCustomerKey(c.ID)
		s.cache.Set(name, 1, 3600)
		s.cache.Set(cKey, c, 3600)
		return nil, errors.New("duplicate customer")
	}

	customer := &domain.Customer{
		Name:      name,
		Status:    "active",
		Balance:   0,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Save(customer); err != nil {
		return nil, err
	}

	// store customer into cache
	if err := s.cache.Set(GetCustomerKey(customer.ID), customer, 0); err != nil {
		return nil, errors.New("failed to store customer into cache")
	}
	return customer, nil
}

func (s *Service) GetCustomer(id uint64) (*domain.Customer, error) {
	// look up in cache first
	var customer *domain.Customer
	var success bool
	var err error
	success, customer = LookUpCustomerInCache(s.cache, id)
	if !success {
		customer, err = s.repo.GetByID(id)
		if err != nil {
			return nil, errors.New("customer not found")
		}

		// store customer into cache
		if err := s.cache.Set(GetCustomerKey(customer.ID), customer, 0); err != nil {
			return nil, errors.New("failed to store customer into cache")
		}
	} else {
		log.Println("get customer from cache!")
	}

	return customer, nil
}

func (s *Service) DeactiveCustomer(id uint64) error {
	customer, err := s.GetCustomer(id)
	if err != nil {
		return err
	}

	if customer.Status != "inactive" {
		customer.Status = "inactive"
		if err := s.cache.Set(GetCustomerKey(customer.ID), customer, 0); err != nil {
			return errors.New("failed to store customer into cache")
		}
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
		if err := s.cache.Set(GetCustomerKey(customer.ID), customer, 0); err != nil {
			return errors.New("failed to store customer into cache")
		}
		return s.repo.Update(customer)
	}

	return nil
}

func (s *Service) SubmitOrder(customerID uint64, items []domain.OrderItem) (bool, error) {
	// look up in cache first
	var customer *domain.Customer
	var success bool
	var err error
	success, customer = LookUpCustomerInCache(s.cache, customerID)
	if !success {
		customer, err = s.repo.GetByID(customerID)
		if err != nil {
			return false, errors.New("customer not found")
		}

		// store customer into cache
		if err := s.cache.Set(GetCustomerKey(customer.ID), customer, 0); err != nil {
			return false, errors.New("failed to store customer into cache")
		}
	}

	if !customer.CanPlaceOrder() {
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

func (s *Service) GetUnpaidOrders(customerID uint64) ([]domain.Order, error) {
	return s.repo.GetUnpaidOrdersByID(customerID)
}

func (s *Service) StoreBalance(customerID uint64, balance float32) error {
	// look up in cache first
	var customer *domain.Customer
	var success bool
	var err error
	success, customer = LookUpCustomerInCache(s.cache, customerID)
	if !success {
		customer, err = s.repo.GetByID(customerID)
		if err != nil {
			return errors.New("customer not found")
		}
	}

	customer.AddBalance(balance)

	// update cache and database
	if err := s.cache.Set(GetCustomerKey(customerID), customer, 0); err != nil {
		return errors.New("failed to update cache")
	}
	if err := s.repo.Update(customer); err != nil {
		return errors.New("failed to updata database")
	}

	return nil
}

func LookUpCustomerInCache(cache ports.Cache, id uint64) (bool, *domain.Customer) {
	var customer *domain.Customer
	if cache.Exists(GetCustomerKey(id)) {
		data, err := cache.Get(GetCustomerKey(id))
		if err != nil {
			return false, nil
		}
		json.Unmarshal(data, &customer)
	} else {
		return false, nil
	}

	return true, customer
}
