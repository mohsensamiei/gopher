package di

import (
	"context"
	"fmt"
)

func Provide[T any](ctx context.Context, key string) T {
	val := ctx.Value(key)
	if val == nil {
		panic(fmt.Sprintf("%v key does not registered in DI", key))
	}
	return val.(T)
}
