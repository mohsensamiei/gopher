package grpcext

import (
	"context"
	"github.com/pinosell/gopher/pkg/errors"
	"github.com/pinosell/gopher/pkg/i18next"
	"google.golang.org/grpc"
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
