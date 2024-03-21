package grpcext

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/di"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/mohsensamiei/gopher/v2/pkg/i18next"
	"github.com/mohsensamiei/gopher/v2/pkg/metadataext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryWrapErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		res, err := handler(ctx, req)
		if err != nil {
			e := errors.Cast(err)
			return nil, e.SetLocalize(i18next.ByContext(ctx, e.Slug()))
		}
		return res, nil
	}
}

func UnaryContextMetadataInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for key, value := range md {
				ctx = metadataext.SetValue(ctx, key, value[0])
			}
		}
		return handler(ctx, req)
	}
}

func DIUnaryInterceptor[T any](tc di.ToContext[T], provide func() T) grpc.UnaryServerInterceptor {
	value := provide()
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		return handler(tc(ctx, value), req)
	}
}
