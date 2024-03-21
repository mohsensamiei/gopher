package jwtext

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/authorize"
	"github.com/mohsensamiei/gopher/pkg/di"
)

func FromContext(ctx context.Context) *JWT {
	return di.Provide[*JWT](ctx, authorize.Name)
}

func ToContext(ctx context.Context, jwt *JWT) context.Context {
	return context.WithValue(ctx, authorize.Name, jwt)
}
