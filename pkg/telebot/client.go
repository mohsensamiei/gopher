package telebot

import "github.com/mohsensamiei/gopher/v3/pkg/telegram"

func New(configs Configs) *Client {
	return &Client{
		channel:  nil,
		Configs:  configs,
		commands: make(map[string]Command),
		events:   make(map[telegram.UpdateType]Event),
	}
}

type Client struct {
	Configs
	channel     chan telegram.Update
	commands    map[string]Command
	middlewares []Middleware
	events      map[telegram.UpdateType]Event
}
