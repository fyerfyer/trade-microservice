package order

import (
	"gorm.io/gorm"
	"trade-microservice.fyerfyer.net/internal/application/domain"
)

type Order struct {
	gorm.Model
	CustomerID uint64      `gorm:"not null;index"`
	Status     string      `gorm:"type:varchar(50);not null"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Cascaded update
}

type OrderItem struct {
	gorm.Model
	ProductCode string  `gorm:"type:varchar(100);not null"`
	UnitPrice   float32 `gorm:"not null"`
	Quantity    int32   `gorm:"not null"`
	OrderID     uint64  `gorm:"not null;index"` // foreign index
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Save(order *domain.Order) error {
	// convert domain object into database object
	orderModel := ConvertDomainOrderIntoDB(*order)

	res := r.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = uint64(orderModel.ID)
	}
	return res.Error
}

func (r *GormRepository) Update(order *domain.Order) error {
	// convert domain object into database object
	orderModel := ConvertDomainOrderIntoDB(*order)
	res := r.db.Save(&orderModel)
	return res.Error
}

func (r *GormRepository) Delete(orderID uint64) error {
	return r.db.Delete(&Order{}, orderID).Error
}

func (r *GormRepository) FindUnpaidByUser(userID uint64) ([]domain.Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ? AND status = ?", userID, "unpaid").
		Find(&orders).
		Error

	if err != nil {
		return nil, err
	}

	// convert database model into domain model
	return ConvertDBIntoDomainOrders(orders), nil
}

func ConvertDomainOrderIntoDB(order domain.Order) Order {
	var orderItems []OrderItem
	for _, orderItem := range order.Items {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	return Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
}

func ConvertDomainOrdersIntoDB(orders []domain.Order) []Order {
	var dbOrders []Order
	for _, order := range orders {
		dbOrders = append(dbOrders, ConvertDomainOrderIntoDB(order))
	}

	return dbOrders
}

func ConvertDBIntoDomainOrder(order Order) domain.Order {
	var orderItems []domain.OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	return domain.Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		Items:      orderItems,
	}
}

func ConvertDBIntoDomainOrders(orders []Order) []domain.Order {
	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, ConvertDBIntoDomainOrder(order))
	}

	return domainOrders
}
