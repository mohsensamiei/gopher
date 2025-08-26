package telebot

import (
	"context"
	"errors"
	"fmt"
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
	conn := di.Provide[*telegram.Connection](ctx)
	if c.TelegramPullInterval > 0 {
		if ok, err := conn.DeleteWebhook(telegram.DeleteWebhook{}); err != nil {
			return err
		} else if !ok {
			return errors.New("can not delete telegram webhook")
		}
		var updateId uint
		ticker := time.NewTicker(c.TelegramPullInterval)
		for range ticker.C {
			updates, err := conn.GetUpdates(telegram.GetUpdates{
				Offset: updateId,
				Limit:  c.TelegramConcurrency,
			})
			if err != nil {
				return err
			}
			for _, update := range updates {
				updateId = update.UpdateID + 1
				c.channel <- update
			}
		}
	} else {
		if ok, err := conn.SetWebhook(telegram.SetWebhook{
			URL:         fmt.Sprint(c.ExternalURL, hookPath),
			SecretToken: c.TelegramSecretToken,
		}); err != nil {
			return err
		} else if !ok {
			return errors.New("can not set telegram webhook")
		}
	}
	return nil
}
