package telebot

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/di"
)

const (
	Name = "telebot"
)

func FromContext(ctx context.Context) *Client {
	return di.Provide[*Client](ctx, Name)
}

func ToContext(ctx context.Context, client *Client) context.Context {
	return di.Register(ctx, Name, client)
}
