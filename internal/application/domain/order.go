package domain

import "time"

type Order struct {
	ID         uint64
	CustomerID uint64
	Items      []OrderItem
	Status     string // "unprocessed" or "unpaid"
	CreatedAt  int64
}

type OrderItem struct {
	ProductCode string
	UnitPrice   float32
	Quantity    int32
}

func NewOrder(customerID uint64, orderItems []OrderItem) *Order {
	return &Order{
		CustomerID: customerID,
		Items:      orderItems,
		Status:     "unprocessed",
		CreatedAt:  time.Now().Unix(),
	}
}

func (o *Order) TotalPrice() float32 {
	var total float32
	for _, item := range o.Items {
		total += item.UnitPrice * float32(item.Quantity)
	}
	return total
}
