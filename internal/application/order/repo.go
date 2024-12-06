package order

import "trade-microservice.fyerfyer.net/internal/application/domain"

type Repository interface {
	Save(order *domain.Order) error
	Update(order *domain.Order) error
	Delete(orderID uint64) error
}
