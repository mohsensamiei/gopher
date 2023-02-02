package di

import (
	"context"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"net/http"
)

func UnaryRegisterInterceptor(key string, val any) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		ctx = context.WithValue(ctx, key, val)
		return handler(ctx, req)
	}
}

func RegisterMiddleware(key string, val any) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx = context.WithValue(ctx, key, val)
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}
