package order

import "fmt"

func GetOrderKey(id uint64) string {
	return fmt.Sprintf("Order_%v", id)
}
