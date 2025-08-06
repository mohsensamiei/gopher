package worker

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
)

type Builder func(context.Context) context.Context

func DIBuilder[T any](provide func() T) Builder {
	dep := provide()
	return func(ctx context.Context) context.Context {
		return di.Register[T](ctx, dep)
	}
}
