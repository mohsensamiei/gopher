package response

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/di"
	cache "github.com/victorspringer/http-cache"
)

const (
	Name = "response"
)

func FromContext(ctx context.Context) *cache.Client {
	return di.Provide[*cache.Client](ctx, Name)
}

func ToContext(ctx context.Context, client *cache.Client) context.Context {
	return context.WithValue(ctx, Name, client)
}
