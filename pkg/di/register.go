package di

import (
	"context"
)

func Register[T any](ctx context.Context, dep T) context.Context {
	key := TypeName[T]()
	return context.WithValue(ctx, key, dep)
}
