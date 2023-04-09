package gormext

import (
	"context"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net/http"
)

func DIMiddleware(db *gorm.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx = context.WithValue(ctx, Name, db.WithContext(ctx))
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}

func UnaryDIInterceptor(db *gorm.DB) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		return handler(
			context.WithValue(ctx, Name, db.WithContext(ctx)),
			req,
		)
	}
}
