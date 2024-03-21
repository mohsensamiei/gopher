package slug

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/di"
)

const (
	Name = "Slug"
)

func FromContext(ctx context.Context) *Service {
	return di.Provide[*Service](ctx, Name)
}

func ToContext(ctx context.Context, v *Service) context.Context {
	return context.WithValue(ctx, Name, v)
}
