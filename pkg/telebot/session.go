package telebot

import (
	"context"
	"fmt"
	"github.com/mohsensamiei/gopher/v2/pkg/redisext"
)

func (c *Client) SetSession(ctx context.Context, chatID int64, data any) error {
	return redisext.FromContext(ctx).Set(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":sessions"), fmt.Sprint(chatID), data, 0)
}

func (c *Client) GetSession(ctx context.Context, chatID int64, data any) error {
	return redisext.FromContext(ctx).Get(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":sessions"), fmt.Sprint(chatID), data)
}

func (c *Client) GetState(ctx context.Context, chatID int64, data any) error {
	return redisext.FromContext(ctx).Get(ctx, fmt.Sprint(c.TelegramStoragePrefix, ":sessions"), fmt.Sprint(chatID), data)
}
