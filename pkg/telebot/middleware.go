package telebot

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/telegram"
)

type Middleware func(ctx context.Context, update telegram.Update) (bool, error)

func (c *Client) Middleware(mdl Middleware) {
	c.Middlewares = append(c.Middlewares, mdl)
}
