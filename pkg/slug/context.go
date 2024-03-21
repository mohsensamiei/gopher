package slug

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/di"
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
