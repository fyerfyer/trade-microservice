package order

import (
	"context"
	"log"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/grpc_config"
	"trade-microservice.fyerfyer.net/internal/tls"
	pb "trade-microservice.fyerfyer.net/proto/order"
)

type OrderAdapter struct {
	order pb.OrderClient
}

var cb *gobreaker.CircuitBreaker

func initCircuitBreaker() {
	cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "orderServiceCircuitBreaker",
		MaxRequests: 3,
		Timeout:     time.Second * 4,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("cb %s change from status %v to %v", name, from, to)
		},
	})
}

func NewOrderAdapter(orderServiceURL string) (*OrderAdapter, error) {
	// tls config
	tlsCredentials, err := tls.GetClientTLSCredentials()
	if err != nil {
		return err
	}

	// config reset strategy
	opts := grpc_config.ResetModel(codes.Internal,
		5, grpc_retry.BackoffLinear(time.Second))

	// append cb
	initCircuitBreaker()
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_config.CBClientInterceptor(cb)))
	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	conn, err := grpc.Dial(orderServiceURL, opts...)

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
