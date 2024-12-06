package order

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	redis "trade-microservice.fyerfyer.net/internal/application/cache"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/e"
	"trade-microservice.fyerfyer.net/internal/ports"
)

type Service struct {
	repo    Repository
	cache   ports.Cache
	payment ports.PaymentPort
}

func NewService(repo Repository, cache ports.Cache, paymentPort ports.PaymentPort) *Service {
	return &Service{
		repo:    repo,
		cache:   cache,
		payment: paymentPort,
	}
}

func (s *Service) ProcessItems(customerID uint64, items []domain.OrderItem) error {
	order := domain.NewOrder(customerID, items)

	err := s.repo.Save(order)
	if err != nil {
		return e.FAILED_TO_STORE_DB_ERROR
	}

	return s.handlePayment(order)
}

func (s *Service) ProcessOrder(order *domain.Order) error {
	return s.handlePayment(order)
}

func (s *Service) handlePayment(order *domain.Order) error {
	res, err := s.payment.Charge(order.CustomerID, order.ID, order.TotalPrice())

	if err == nil && res.Status == "success" {
		order.Status = "success"
		if err := s.cache.Set(redis.GetOrderKey(order.ID, order.Status), order, 0); err != nil {
			return e.FAILED_TO_UPDATE_CACHE_ERROR
		}
		if err := s.repo.Update(order); err != nil {
			return e.FAILED_TO_UPDATE_DB_ERROR
		}

		return nil
	}

	order.Status = "unpaid"
	if err := s.cache.Set(redis.GetOrderKey(order.ID, order.Status), order, 0); err != nil {
		return e.FAILED_TO_UPDATE_CACHE_ERROR
	}
	if err := s.repo.Update(order); err != nil {
		return e.FAILED_TO_UPDATE_DB_ERROR
	}

	return status.New(codes.Internal, res.Message).Err()
}
