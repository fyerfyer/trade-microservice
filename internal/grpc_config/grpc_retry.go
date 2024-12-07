package grpc_config

import (
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func ResetModel(codes codes.Code, maxRetry uint, backoff grpc_retry.BackoffFunc) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes),
			grpc_retry.WithMax(maxRetry),
			grpc_retry.WithBackoff(backoff),
		)),
	}
}
