package cache

import (
	"encoding/json"
	"fmt"

	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/ports"
)

func LookUpOrderInCache(cache ports.Cache, id uint64, status string) (bool, *domain.Order) {
	var order *domain.Order
	if cache.Exists(GetOrderKey(id, status)) {
		data, err := cache.Get(GetOrderKey(id, status))
		if err != nil {
			return false, nil
		}
		json.Unmarshal(data, &order)
	} else {
		return false, nil
	}

	return true, order
}

func LookUpCustomerInCache(cache ports.Cache, id uint64) (bool, *domain.Customer) {
	var customer *domain.Customer
	if cache.Exists(GetCustomerKey(id)) {
		data, err := cache.Get(GetCustomerKey(id))
		if err != nil {
			return false, nil
		}
		json.Unmarshal(data, &customer)
	} else {
		return false, nil
	}

	return true, customer
}

func GetCustomerKey(id uint64) string {
	return fmt.Sprintf("Customer_%v", id)
}

func GetOrderKey(id uint64, status string) string {
	return fmt.Sprintf("Order_%v_%v", id, status)
}
