package customer

import (
	"gorm.io/driver/mysql"
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

func NewGormRepository(dsn string) (*GormRepository, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Customer{}); err != nil {
		return nil, err
	}
	return &GormRepository{db: db}, nil
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
	err := r.db.Where("id = ?", customerID).First(&customer).Error
	if err != nil {
		return nil, err
	}

	domainCustomer := ConvertDBIntoDomainCustomer(customer)
	return &domainCustomer, nil
}

func (r *GormRepository) GetByName(customerName string) (*domain.Customer, error) {
	var customer Customer
	err := r.db.Where("name = ?", customerName).First(&customer).Error
	if err != nil {
		return nil, err
	}

	domainCustomer := ConvertDBIntoDomainCustomer(customer)
	return &domainCustomer, nil
}

func (r *GormRepository) GetUnpaidOrdersByID(customerID uint64) ([]domain.Order, error) {
	var orders []order.Order
	err := r.db.Preload("OrderItems").Where("status = ? and customer_id = ?", "unpaid", customerID).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	domainOrders := order.ConvertDBIntoDomainOrders(orders)
	return domainOrders, nil
}

func (r *GormRepository) GetOrderByID(customerID, orderID uint64) (*domain.Order, error) {
	var o order.Order
	err := r.db.Preload("OrderItems").Where("customer_id = ? and id = ?", customerID, orderID).First(&o).Error
	if err != nil {
		return nil, err
	}

	domainOrder := order.ConvertDBIntoDomainOrder(o)
	return &domainOrder, nil
}

func ConvertDomainCustomerIntoDB(customer domain.Customer) Customer {
	return Customer{
		Model: gorm.Model{
			ID:        uint(customer.ID),
			CreatedAt: customer.CreatedAt,
		},
		Name:    customer.Name,
		Status:  customer.Status,
		Balance: customer.Balance,
	}
}

func ConvertDBIntoDomainCustomer(customer Customer) domain.Customer {
	return domain.Customer{
		ID:        uint64(customer.ID),
		Name:      customer.Name,
		Status:    customer.Status,
		Balance:   customer.Balance,
		CreatedAt: customer.CreatedAt,
	}
}
