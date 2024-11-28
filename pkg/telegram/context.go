package telegram

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/di"
)

const (
	Name = "telegram"
)

func FromContext(ctx context.Context) *Connection {
	return di.Provide[*Connection](ctx, Name)
}

func ToContext(ctx context.Context, conn *Connection) context.Context {
	return context.WithValue(ctx, Name, conn)
}
