package domain

type Customer struct {
	ID      uint64
	Name    string
	Status  string // "active" or "inactive"
	Balance float32
}

func (c *Customer) CanPlaceOrder() bool {
	return c.Status == "active"
}

func (c *Customer) DeductBalance(amount float32) bool {
	if c.Balance >= amount {
		c.Balance -= amount
		return true
	}
	return false
}
