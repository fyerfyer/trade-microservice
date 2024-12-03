package payment

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	pb "trade-microservice.fyerfyer.net/proto/payment"
)

// the adapter helps order service to dial for payment grpc service
type PaymentAdapter struct {
	payment pb.PaymentClient
}

func NewPaymentAdapter(paymentServiceURL string) (*PaymentAdapter, error) {
	conn, err := grpc.Dial(paymentServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return &PaymentAdapter{payment: pb.NewPaymentClient(conn)}, nil
}

func (p *PaymentAdapter) Charge(customerID uint64, orderID uint64, totalPrice float32) (*domain.Payment, error) {
	res, err := p.payment.Charge(context.Background(), &pb.ChargeRequest{
		CustomerId: customerID,
		OrderId:    orderID,
		TotalPrice: totalPrice,
	})

	if err != nil {
		return nil, err
	}

	return &domain.Payment{
		Status:  res.Status,
		Message: res.Message,
	}, nil
}
