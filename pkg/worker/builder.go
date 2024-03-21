package worker

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/di"
)

type Builder func(context.Context) context.Context

func DIBuilder[T any](tc di.ToContext[T], provide func() T) Builder {
	value := provide()
	return func(ctx context.Context) context.Context {
		return tc(ctx, value)
	}
}
