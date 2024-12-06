package order

import (
	"context"

	"trade-microservice.fyerfyer.net/internal/application/domain"
	pb "trade-microservice.fyerfyer.net/proto/order"
)

func (a *Adapter) ProcessItems(ctx context.Context, req *pb.ProcessItemsRequest) (*pb.ProcessItemsResponse, error) {
	// convert grpc object into domain object
	var items []domain.OrderItem
	for _, item := range req.OrderItems {
		items = append(items, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	err := a.service.ProcessItems(req.GetCustomerId(), items)
	if err != nil {
		return &pb.ProcessItemsResponse{Message: err.Error()}, err
	}

	return &pb.ProcessItemsResponse{Message: "successfully pay for the items"}, nil
}

func (a *Adapter) ProcessOrder(ctx context.Context, req *pb.ProcessOrderRequest) (*pb.ProcessOrderResponse, error) {
	// convert grpc object into domain object
	var items []domain.OrderItem
	for _, item := range req.Order.OrderItems {
		items = append(items, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	domainOrder := &domain.Order{
		ID:     req.GetOrder().OrderId,
		Items:  items,
		Status: req.GetOrder().Status,
	}

	err := a.service.ProcessOrder(domainOrder)
	if err != nil {
		return &pb.ProcessOrderResponse{Message: err.Error()}, err
	}

	return &pb.ProcessOrderResponse{Message: "successfully purchase the order"}, nil
}
