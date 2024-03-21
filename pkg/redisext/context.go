package redisext

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/di"
)

const (
	Name = "Redis"
)

func FromContext(ctx context.Context) *Client {
	return di.Provide[*Client](ctx, Name)
}

func ToContext(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, Name, client)
}
