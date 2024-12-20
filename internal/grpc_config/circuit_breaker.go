package grpc_config

import (
	"context"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
)

func CBClientInterceptor(cb *gobreaker.CircuitBreaker) grpc.UnaryClientInterceptor {
	return func(ctx context.Context,
		method string,
		req, res interface{},
		conn *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		_, cbErr := cb.Execute(func() (interface{}, error) {
			err := invoker(ctx, method, req, res, conn, opts...)
			if err != nil {
				return nil, err
			}

			return nil, nil
		})
		return cbErr
	}
}
