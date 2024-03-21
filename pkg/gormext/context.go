package gormext

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/di"
	"gorm.io/gorm"
)

func FromContext(ctx context.Context) *gorm.DB {
	return di.Provide[*gorm.DB](ctx, Name).WithContext(ctx)
}

func ToContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, Name, db)
}
