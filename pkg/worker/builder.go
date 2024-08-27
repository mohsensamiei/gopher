package worker

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/di"
)

type Builder func(context.Context) context.Context

func DIBuilder[T any](key any, provide func() T) Builder {
	value := provide()
	return func(ctx context.Context) context.Context {
		return di.Register(ctx, key, value)
	}
}
