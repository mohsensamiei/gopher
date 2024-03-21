package telebot

import "github.com/mohsensamiei/gopher/pkg/telegram"

func New(configs Configs, connection *telegram.Connection) *Client {
	return &Client{
		channel:    nil,
		Configs:    configs,
		Connection: connection,
		commands:   make(map[string]Command),
	}
}

type Client struct {
	Configs
	Commands    []Command
	Middlewares []Middleware
	*telegram.Connection
	channel  chan telegram.Update
	commands map[string]Command
}
