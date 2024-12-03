package order

import (
	"context"

	"trade-microservice.fyerfyer.net/internal/application/domain"
	pb "trade-microservice.fyerfyer.net/proto/order"
)

func (a *Adapter) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// convert grpc object into domain object
	var items []domain.OrderItem
	for _, item := range req.OrderItems {
		items = append(items, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	res, err := a.api.CreateOrder(req.GetCustomerId(), items)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{OrderId: res.ID}, nil
}

func (a *Adapter) GetUnpaidOrders(ctx context.Context, req *pb.GetUnpaidOrdersRequest) (*pb.GetUnpaidOrdersResponse, error) {
	orders, err := a.api.GetUnpaidOrders(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	var grpcOrders []*pb.OrderEntity
	for _, order := range orders {
		var items []*pb.OrderItem
		for _, item := range order.Items {
			items = append(items, &pb.OrderItem{
				ProductCode: item.ProductCode,
				UnitPrice:   item.UnitPrice,
				Quantity:    item.Quantity,
			})
		}

		grpcOrders = append(grpcOrders, &pb.OrderEntity{
			OrderId: order.ID,
			Items:   items,
			Status:  order.Status,
		})
	}

	return &pb.GetUnpaidOrdersResponse{UnpaidOrders: grpcOrders}, nil
}
