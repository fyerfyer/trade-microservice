package payment

import (
	"context"

	pb "trade-microservice.fyerfyer.net/proto/payment"
)

func (a *Adapter) Charge(ctx context.Context, req *pb.ChargeRequest) (*pb.ChargeResponse, error) {
	payment, err := a.api.Charge(req.GetCustomerId(), req.GetOrderId(), req.GetTotalPrice())
	if err != nil {
		return nil, err
	}

	return &pb.ChargeResponse{
		Status:  payment.Status,
		Message: payment.Message,
	}, nil
}
