package gormext

import (
	"context"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

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
