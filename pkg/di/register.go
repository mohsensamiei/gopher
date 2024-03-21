package di

import (
	"context"
)

func Register(ctx context.Context, key string, val any) context.Context {
	return context.WithValue(ctx, key, val)
}
