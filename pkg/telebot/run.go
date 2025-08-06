package telebot

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"github.com/mohsensamiei/gopher/v3/pkg/telegram"
	log "github.com/sirupsen/logrus"
	"time"
)

func (c *Client) Run(ctx context.Context) error {
	c.channel = make(chan telegram.Update, c.TelegramConcurrency)
	defer func() {
		close(c.channel)
	}()
	for i := uint8(0); i < c.TelegramConcurrency; i++ {
		go func() {
			for update := range c.channel {
				if err := c.recoveredProcess(ctx, update); err != nil {
					log.WithError(err).Error("can not process message")
				}
			}
		}()
	}
	var updateID uint
	ticker := time.NewTicker(c.TelegramPullInterval)
	for range ticker.C {
		updates, err := di.Provide[*telegram.Connection](ctx).GetUpdates(telegram.GetUpdates{
			Offset: updateID,
			Limit:  c.TelegramConcurrency,
		})
		if err != nil {
			return err
		}
		for _, update := range updates {
			updateID = update.UpdateID + 1
			c.channel <- update
		}
	}
	return nil
}
