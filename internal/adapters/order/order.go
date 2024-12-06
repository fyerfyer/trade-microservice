package order

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	pb "trade-microservice.fyerfyer.net/proto/order"
)

type OrderAdapter struct {
	order pb.OrderClient
}

func NewOrderAdapter(orderServiceURL string) (*OrderAdapter, error) {
	conn, err := grpc.Dial(orderServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return &OrderAdapter{order: pb.NewOrderClient(conn)}, nil
}

func (o *OrderAdapter) ProcessItems(customerID uint64, items []domain.OrderItem) error {
	_, err := o.order.ProcessItems(context.Background(), &pb.ProcessItemsRequest{
		CustomerId: customerID,
		OrderItems: convertDomainItemsIntoPB(items),
	})

	if err != nil {
		return err
	}

	return nil
}

func (o *OrderAdapter) ProcessOrder(order domain.Order) error {
	_, err := o.order.ProcessOrder(context.Background(), &pb.ProcessOrderRequest{
		Order: &pb.OrderEntity{
			OrderId:    order.ID,
			OrderItems: convertDomainItemsIntoPB(order.Items),
			Status:     order.Status,
		},
	})

	if err != nil {
		return err
	}

	return nil
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
