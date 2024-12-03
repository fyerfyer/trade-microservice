package order

import (
	"context"

	"trade-microservice.fyerfyer.net/internal/application/domain"
	pb "trade-microservice.fyerfyer.net/proto/order"
)

func (a *Adapter) CreateOrder(ctx context.Context, req *pb.ProcessOrderRequest) (*pb.ProcessOrderResponse, error) {
	// convert grpc object into domain object
	var items []domain.OrderItem
	for _, item := range req.OrderItems {
		items = append(items, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	err := a.api.ProcessOrder(req.GetCustomerId(), items)
	if err != nil {
		return &pb.ProcessOrderResponse{Status: "failure"}, err
	}

	return &pb.ProcessOrderResponse{Status: "success"}, nil
}
