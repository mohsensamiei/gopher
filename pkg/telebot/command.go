package telebot

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/telegram"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Action func(ctx context.Context, update telegram.Update, data any) (string, any, error)

type Route struct {
	Action       Action
	AllowUpdates []telegram.UpdateType
}

type Command interface {
	Name() string
	Alias() []string
	Description() string
	Init(ctx context.Context, update telegram.Update, data any) (string, any, error)
	Actions() map[string]Route
}

func (c *Client) Command(cmd Command) {
	names := append(cmd.Alias(), cmd.Name())
	for _, name := range names {
		name = strings.ToLower(strings.TrimSpace(name))
		if _, ok := c.commands[name]; ok {
			log.WithField("command", cmd.Name()).Fatal("duplicate command name")
		}
		c.commands[name] = cmd
	}
	c.Commands = append(c.Commands, cmd)
}
