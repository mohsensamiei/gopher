package telebot

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/telegram"
)

type Event func(ctx context.Context, update telegram.Update) error

func (c *Client) Event(on telegram.UpdateType, event Event) {
	c.events[on] = event
}
