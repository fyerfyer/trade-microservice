package customer

import (
	"context"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	pb "trade-microservice.fyerfyer.net/proto/customer"
)

func (a *Adapter) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	customer, err := a.service.CreateCustomer(req.GetCustomerName())
	if err != nil {
		return nil, err
	}

	return &pb.CreateCustomerResponse{
		CustomerId: customer.ID,
		Success:    true,
	}, nil
}

func (a *Adapter) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customer, err := a.service.GetCustomer(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.GetCustomerResponse{Customer: &pb.CustomerEntity{
		CustomerId: customer.ID,
		Name:       customer.Name,
		Status:     customer.Status,
		Balance:    customer.Balance,
	}}, nil
}

func (a *Adapter) DeactiveCustomer(ctx context.Context, req *pb.DeactivateCustomerRequest) (*pb.DeactivateCustomerResponse, error) {
	err := a.service.DeactiveCustomer(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.DeactivateCustomerResponse{Status: "inactive"}, nil
}

func (a *Adapter) ActivateCustomer(ctx context.Context, req *pb.ActivateCustomerRequest) (*pb.ActivateCustomerResponse, error) {
	err := a.service.ReactiveCustomer(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.ActivateCustomerResponse{Status: "active"}, nil
}

func (a *Adapter) SubmitOrder(ctx context.Context, req *pb.SubmitOrderRequest) (*pb.SubmitOrderResponse, error) {
	success, err := a.service.SubmitOrder(req.GetCustomerId(), convertPBIntoDomainItems(req.GetOrderItems()))
	if err != nil {
		return &pb.SubmitOrderResponse{
			Success: success,
			Message: err.Error(),
		}, err
	}

	return &pb.SubmitOrderResponse{
		Success: success,
		Message: "items processed successfully",
	}, nil
}

func (a *Adapter) PayOrder(ctx context.Context, req *pb.PayOrderRequest) (*pb.PayOrderResponse, error) {
	success, err := a.service.PayOrder(req.GetCustomerId(), req.GetOrderId())
	if err != nil {
		return &pb.PayOrderResponse{
			Success: success,
			Message: err.Error(),
		}, err
	}

	return &pb.PayOrderResponse{
		Success: success,
		Message: "order processed successfully",
	}, nil
}

func (a *Adapter) GetUnpaidOrders(ctx context.Context, req *pb.GetUnpaidOrdersRequest) (*pb.GetUnpaidOrdersResponse, error) {
	orders, err := a.service.GetUnpaidOrders(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUnpaidOrdersResponse{
		UnpaidOrders: convertDomainOrderIntoPB(orders),
	}, nil
}

func (a *Adapter) StoreBalance(ctx context.Context, req *pb.StoreBalanceRequest) (*pb.StoreBalanceResponse, error) {
	if err := a.service.StoreBalance(req.GetCustomerId(), req.GetBalance()); err != nil {
		return &pb.StoreBalanceResponse{
			Message: err.Error(),
		}, err
	}

	return &pb.StoreBalanceResponse{
		Message: "successfully stored balance",
	}, nil
}

func convertPBIntoDomainItems(items []*pb.OrderItem) []domain.OrderItem {
	var domainItems []domain.OrderItem
	for _, item := range items {
		domainItems = append(domainItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	return domainItems
}

func convertDomainItemsIntoPB(items []domain.OrderItem) []*pb.OrderItem {
	var pbItems []*pb.OrderItem
	for _, item := range items {
		pbItems = append(pbItems, &pb.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	return pbItems
}

func convertDomainOrderIntoPB(orders []domain.Order) []*pb.Order {
	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			OrderId: order.ID,
			Items:   convertDomainItemsIntoPB(order.Items),
			Status:  order.Status,
		})
	}

	return pbOrders
}
