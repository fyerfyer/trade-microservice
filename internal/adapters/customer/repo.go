package customer

import (
	"gorm.io/gorm"
	"trade-microservice.fyerfyer.net/internal/adapters/order"
	"trade-microservice.fyerfyer.net/internal/application/domain"
)

type Customer struct {
	gorm.Model
	Name    string        `gorm:"type:varchar(100);not null"`
	Status  string        `gorm:"type:varchar(50);not null"`
	Balance float32       `gorm:"not null"`
	Orders  []order.Order `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(customer *domain.Customer) error {
	dbCustomer := ConvertDomainCustomerIntoDB(*customer)
	err := r.db.Create(&dbCustomer).Error
	if err == nil {
		customer.ID = uint64(dbCustomer.ID)
	}

	return err
}

func (r *GormRepository) Update(customer *domain.Customer) error {
	dbCustomer := ConvertDomainCustomerIntoDB(*customer)
	return r.db.Save(&dbCustomer).Error
}

func (r *GormRepository) GetByID(customerID uint64) (*domain.Customer, error) {
	var customer Customer
	err := r.db.Where("customer_id = ?", customerID).First(&customer).Error
	if err != nil {
		return nil, err
	}

	domainCustomer := ConvertDBIntoDomainCustomer(customer)
	return &domainCustomer, nil
}

func (r *GormRepository) GetUnpaidOrdersByID(customerID uint64) ([]domain.Order, error) {
	var orders []order.Order
	err := r.db.Where("customer_id = ?", customerID).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	domainOrders := order.ConvertDBIntoDomainOrders(orders)
	return domainOrders, nil
}

func ConvertDomainCustomerIntoDB(customer domain.Customer) Customer {
	return Customer{
		Name:    customer.Name,
		Status:  customer.Status,
		Balance: customer.Balance,
	}
}

func ConvertDBIntoDomainCustomer(customer Customer) domain.Customer {
	return domain.Customer{
		Name:    customer.Name,
		Status:  customer.Status,
		Balance: customer.Balance,
	}
}
