package redisext

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pinosell/gopher/pkg/di"
	"google.golang.org/grpc"
	"net/http"
)

const (
	Name = "Redis"
)

func Provide(ctx context.Context) *Client {
	return di.Provide[*Client](ctx, Name)
}

func UnaryDIInterceptor(client *Client) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		return handler(
			context.WithValue(ctx, Name, client),
			req,
		)
	}
}

func DIMiddleware(client *Client) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx = context.WithValue(ctx, Name, client)
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}
