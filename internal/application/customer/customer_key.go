package customer

import (
	"fmt"
)

func GetCustomerKey(id uint64) string {
	return fmt.Sprintf("Customer_%v", id)
}
