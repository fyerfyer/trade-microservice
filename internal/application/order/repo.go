package order

import "trade-microservice.fyerfyer.net/internal/application/domain"

type Repository interface {
	Save(order *domain.Order) error
	Update(order *domain.Order) error
	Delete(orderID uint64) error
	FindUnpaidByUser(userID uint64) ([]domain.Order, error)
}
