package metadataext

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryContextInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for key, value := range md {
				ctx = SetValue(ctx, key, value[0])
			}
		}
		return handler(ctx, req)
	}
}
