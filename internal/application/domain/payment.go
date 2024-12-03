package domain

import "time"

type Payment struct {
	CustomerID uint64
	OrderID    uint64
	TotalPrice float32
	Status     string // "success" or "failure"
	Message    string
	CreatedAt  time.Time
}
