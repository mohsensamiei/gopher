package telebot

import (
	"context"
	"fmt"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"github.com/mohsensamiei/gopher/v3/pkg/redisext"
)

func (c *Client) SetSession(ctx context.Context, chatID int64, data any) error {
	return di.Provide[*redisext.Client](ctx).Set(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":sessions"), fmt.Sprint(chatID), data, 0)
}

func (c *Client) GetSession(ctx context.Context, chatID int64, data any) error {
	return di.Provide[*redisext.Client](ctx).Get(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":sessions"), fmt.Sprint(chatID), data)
}

func (c *Client) DelSession(ctx context.Context, chatID int64) error {
	return di.Provide[*redisext.Client](ctx).Del(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":sessions"), fmt.Sprint(chatID))
}

func (c *Client) GetState(ctx context.Context, chatID int64, state *State) error {
	return di.Provide[*redisext.Client](ctx).Get(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":states"), fmt.Sprint(chatID), state)
}

func (c *Client) SetState(ctx context.Context, chatID int64, state *State) error {
	return di.Provide[*redisext.Client](ctx).Set(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":states"), fmt.Sprint(chatID), state, 0)
}

func (c *Client) DelState(ctx context.Context, chatID int64) error {
	return di.Provide[*redisext.Client](ctx).Del(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":states"), fmt.Sprint(chatID))
}
