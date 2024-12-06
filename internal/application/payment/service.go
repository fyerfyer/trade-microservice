package payment

import (
	redis "trade-microservice.fyerfyer.net/internal/application/cache"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/e"
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
	var (
		c       *domain.Customer
		success bool
		err     error
	)

	success, c = redis.LookUpCustomerInCache(s.cache, customerID)
	if !success {
		c, err = s.repo.GetCustomerByID(customerID)
		if err != nil {
			return &domain.Payment{
				Status:  "failure",
				Message: "customer not found",
			}, e.CUSTOMER_NOT_FOUND
		}
	}

	if c == nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "customer not found",
		}, e.CUSTOMER_NOT_FOUND
	}

	// purchase
	if !c.DeductBalance(totalAmount) {
		return &domain.Payment{
			Status:  "failure",
			Message: "customer didn't have enough money",
		}, e.BALANCE_INSUFFICIENT
	}

	// to ensure data consistency, update customer cache manually
	if err := s.cache.Set(redis.GetCustomerKey(customerID), c, 0); err != nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "failed to update cache",
		}, e.FAILED_TO_UPDATE_CACHE_ERROR
	}

	// update customer database
	if err := s.repo.UpdateCustomerBalance(c); err != nil {
		return &domain.Payment{
			Status:  "failure",
			Message: "failed to update database",
		}, e.FAILED_TO_UPDATE_DB_ERROR
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
		}, e.FAILED_TO_STORE_DB_ERROR
	}

	return payment, nil
}
