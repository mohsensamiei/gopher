package di

import (
	"context"
	"fmt"
)

func Provide[T any](ctx context.Context) T {
	key := TypeName[T]()
	val := ctx.Value(key)
	if val == nil {
		panic(fmt.Sprintf("%v does not registered in di", key))
	}
	return val.(T)
}
