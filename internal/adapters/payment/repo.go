package payment

import (
	"gorm.io/gorm"
	"trade-microservice.fyerfyer.net/internal/adapters/customer"
	"trade-microservice.fyerfyer.net/internal/application/domain"
)

type Payment struct {
	CustomerID uint64  `gorm:"not null;index"`
	OrderID    uint64  `gorm:"primaryKey;not null;index"`
	TotalPrice float32 `gorm:"not null"`
	Status     string  `gorm:"type:varchar(50);not null"`
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) GetCustomerByID(customerID uint64) (*domain.Customer, error) {
	var c customer.Customer
	if err := r.db.Model(&customer.Customer{}).Where("id = ?", customerID).First(&c, customerID).Error; err != nil {
		return nil, err
	}

	domainCustomer := customer.ConvertDBIntoDomainCustomer(c)
	return &domainCustomer, nil
}

func (r *GormRepository) UpdateCustomerBalance(c *domain.Customer) error {
	// convert domain object into db object
	dbCustomer := customer.ConvertDomainCustomerIntoDB(*c)
	return r.db.Model(&customer.Customer{}).Where("id = ?", c.ID).Save(&dbCustomer).Error
}

func (r *GormRepository) SavePayment(payment *domain.Payment) error {
	dbPayment := ConvertDomainPaymentIntoDB(*payment)
	return r.db.Create(&dbPayment).Error
}

func ConvertDomainPaymentIntoDB(payment domain.Payment) Payment {
	return Payment{
		CustomerID: payment.CustomerID,
		OrderID:    payment.OrderID,
		TotalPrice: payment.TotalPrice,
		Status:     payment.Status,
	}
}

// func ConvertDBIntoDomainPayment(payment Payment) domain.Payment {
// 	return domain.Payment{
// 		CustomerID: payment.CustomerID,
// 		OrderID:    payment.OrderID,
// 		TotalPrice: payment.TotalPrice,
// 		Status:     payment.Status,
// 	}
// }
