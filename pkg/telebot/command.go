package telebot

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/stringsext"
	"github.com/mohsensamiei/gopher/v2/pkg/telegram"
	log "github.com/sirupsen/logrus"
)

type Keyword string

func (k Keyword) String() string {
	return stringsext.Comparable(string(k))
}

const (
	Empty Keyword = ""
)

type Action func(ctx context.Context, update telegram.Update) (Keyword, error)

type Route struct {
	Action       Action
	AllowUpdates []telegram.UpdateType
}

type Command interface {
	Name() Keyword
	Alias() []Keyword
	Actions() map[Keyword]Route
	Init(ctx context.Context, update telegram.Update) (Keyword, error)
}

func (c *Client) Command(cmd Command) {
	keywords := append(cmd.Alias(), cmd.Name())
	for _, keyword := range keywords {
		if _, ok := c.commands[keyword.String()]; ok {
			log.WithField("command", cmd.Name()).Fatal("duplicate command name")
		}
		c.commands[keyword.String()] = cmd
	}
}
