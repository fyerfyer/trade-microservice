package customer

import (
	"time"

	redis "trade-microservice.fyerfyer.net/internal/application/cache"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/e"
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
		return nil, e.DUPLICATE_CUSTOMER_ERROR
	}

	// else, find in db
	// if we can find in db, then update our cache
	// we store id and name key separately
	if c, _ := s.repo.GetByName(name); c != nil {
		cKey := redis.GetCustomerKey(c.ID)
		s.cache.Set(name, 1, 0)
		s.cache.Set(cKey, c, 0)
		return nil, e.DUPLICATE_CUSTOMER_ERROR
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
	if err := s.cache.Set(redis.GetCustomerKey(customer.ID), customer, 0); err != nil {
		return nil, e.FAILED_TO_STORE_CACHE_ERROR
	}
	return customer, nil
}

func (s *Service) GetCustomer(id uint64) (*domain.Customer, error) {
	// look up in cache first
	var (
		customer *domain.Customer
		success  bool
		err      error
	)
	success, customer = redis.LookUpCustomerInCache(s.cache, id)
	if !success {
		customer, err = s.repo.GetByID(id)
		if err != nil {
			return nil, e.CUSTOMER_NOT_FOUND
		}

		// store customer into cache
		if err := s.cache.Set(redis.GetCustomerKey(customer.ID), customer, 0); err != nil {
			return nil, e.FAILED_TO_STORE_CACHE_ERROR
		}
	}

	return customer, nil
}

func (s *Service) DeactiveCustomer(id uint64) error {
	customer, err := s.GetCustomer(id)
	if err != nil {
		return e.CUSTOMER_NOT_FOUND
	}

	if customer.Status != "inactive" {
		customer.Status = "inactive"
		if err := s.cache.Set(redis.GetCustomerKey(customer.ID), customer, 0); err != nil {
			return e.FAILED_TO_UPDATE_CACHE_ERROR
		}
		if err := s.repo.Update(customer); err != nil {
			return e.FAILED_TO_UPDATE_DB_ERROR
		}
	}

	return nil
}

func (s *Service) ReactiveCustomer(id uint64) error {
	customer, err := s.GetCustomer(id)
	if err != nil {
		return e.CUSTOMER_NOT_FOUND
	}

	if customer.Status != "active" {
		customer.Status = "active"
		if err := s.cache.Set(redis.GetCustomerKey(customer.ID), customer, 0); err != nil {
			return e.FAILED_TO_STORE_CACHE_ERROR
		}
		if err := s.repo.Update(customer); err != nil {
			return e.FAILED_TO_UPDATE_DB_ERROR
		}
	}

	return nil
}

func (s *Service) SubmitOrder(customerID uint64, items []domain.OrderItem) (bool, error) {
	// look up in cache first
	var (
		customer *domain.Customer
		success  bool
		err      error
	)

	success, customer = redis.LookUpCustomerInCache(s.cache, customerID)
	if !success {
		customer, err = s.repo.GetByID(customerID)
		if err != nil {
			return false, e.CUSTOMER_NOT_FOUND
		}

		// store customer into cache
		if err := s.cache.Set(redis.GetCustomerKey(customer.ID), customer, 0); err != nil {
			return false, e.FAILED_TO_STORE_CACHE_ERROR
		}
	}

	if !customer.CanPlaceOrder() {
		return false, e.CUSTOMER_INACTIVE
	}

	// the CreateOrderMethod will process the order
	err = s.order.ProcessItems(customerID, items)

	// failed to pay
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Service) PayOrder(customerID, orderID uint64) (bool, error) {
	var (
		customer *domain.Customer
		order    *domain.Order
		success  bool
		err      error
	)

	success, customer = redis.LookUpCustomerInCache(s.cache, customerID)
	if !success {
		customer, err = s.repo.GetByID(customerID)
		if err != nil {
			return false, e.CUSTOMER_NOT_FOUND
		}

		// store customer into cache
		if err := s.cache.Set(redis.GetCustomerKey(customer.ID), customer, 0); err != nil {
			return false, e.FAILED_TO_STORE_CACHE_ERROR
		}
	}

	if !customer.CanPlaceOrder() {
		return false, e.CUSTOMER_INACTIVE
	}

	success, order = redis.LookUpOrderInCache(s.cache, orderID, "unpaid")
	if !success {
		order, err = s.repo.GetOrderByID(customerID, orderID)
		if err != nil {
			return false, e.ORDER_NOT_FOUND
		}

		if order.Status != "unpaid" {
			return false, e.ORDER_INVALID
		}
	}

	err = s.order.ProcessOrder(order)
	if err != nil {
		return false, err
	}

	// if success, then the order is no longer unpaid
	// so we remove the unpaid key from redis
	if _, err := s.cache.Delete(redis.GetOrderKey(orderID, "unpaid")); err != nil {
		return true, e.FAILED_TO_UPDATE_CACHE_ERROR
	}
	return true, nil
}

func (s *Service) GetUnpaidOrders(customerID uint64) ([]domain.Order, error) {
	order, err := s.repo.GetUnpaidOrdersByID(customerID)
	if err != nil {
		return nil, e.ORDER_NOT_FOUND
	}

	return order, nil
}

func (s *Service) StoreBalance(customerID uint64, balance float32) error {
	// look up in cache first
	var (
		customer *domain.Customer
		success  bool
		err      error
	)
	success, customer = redis.LookUpCustomerInCache(s.cache, customerID)
	if !success {
		customer, err = s.repo.GetByID(customerID)
		if err != nil {
			return e.CUSTOMER_NOT_FOUND
		}
	}

	customer.AddBalance(balance)

	// update cache and database
	if err := s.cache.Set(redis.GetCustomerKey(customerID), customer, 0); err != nil {
		return e.FAILED_TO_UPDATE_CACHE_ERROR
	}
	if err := s.repo.Update(customer); err != nil {
		return e.FAILED_TO_UPDATE_DB_ERROR
	}

	return nil
}
