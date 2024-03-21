package s3ext

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/di"
)

const (
	Name = "s3"
)

func FromContext(ctx context.Context) *Client {
	return di.Provide[*Client](ctx, Name)
}

func ToContext(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, Name, client)
}
