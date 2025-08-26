package telebot

import (
	"context"
	"fmt"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/mapext"
	"github.com/mohsensamiei/gopher/v3/pkg/redisext"
	"github.com/mohsensamiei/gopher/v3/pkg/slices"
	"github.com/mohsensamiei/gopher/v3/pkg/telegram"
	"google.golang.org/grpc/codes"
	"runtime/debug"
	"strings"
	"time"
)

func parseCommand(str string) (Keyword, []string) {
	dump := strings.Split(str, " ")
	command := strings.ToLower(strings.TrimSpace(dump[0]))
	if len(dump) > 1 {
		return Keyword(command), strings.Split(dump[1], "_")
	}
	return Keyword(command), nil
}

type State struct {
	Command   Keyword  `json:"command,omitempty"`
	Action    Keyword  `json:"action,omitempty"`
	Arguments []string `json:"arguments,omitempty"`
}

func (c *Client) recoveredProcess(ctx context.Context, update telegram.Update) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%s:\n%s", rec, debug.Stack())
		}
	}()
	return c.process(ctx, update)
}

func (c *Client) process(ctx context.Context, update telegram.Update) error {
	if update.Chat() == nil {
		return errors.New(codes.Unimplemented).WithDetailF("unsupported update type '%v'", update.Type())
	}
	{
		mutex := di.Provide[*redisext.Client](ctx).NewMutex(fmt.Sprint(c.TelegramStoragePrefix, ":locks"), fmt.Sprint(update.Chat().ID), time.Minute)
		if err := mutex.LockContext(ctx); err != nil {
			return err
		}
		defer func() {
			_, _ = mutex.UnlockContext(ctx)
		}()
	}
	if event, ok := c.events[update.Type()]; ok {
		action := func(ctx context.Context, update telegram.Update) (Keyword, error) {
			if err := event(ctx, update); err != nil {
				return Empty, err
			}
			return Empty, nil
		}
		for _, middleware := range c.middlewares {
			action = middleware(action)
		}
		return event(ctx, update)
	}
	{
		var state = new(State)
		switch err := c.GetState(ctx, update.Chat().ID, state); errors.Code(err) {
		case codes.OK, codes.NotFound:
		case codes.InvalidArgument:
			_ = c.DelState(ctx, update.Chat().ID)
		default:
			return err
		}
		var action Action
		{
			if update.IsCommand() || update.IsSimilarCommand(mapext.Keys(c.commands)) {
				state.Command, state.Arguments = parseCommand(update.Message.Text)
				if err := c.SetState(ctx, update.Chat().ID, state); err != nil {
					return err
				}
				if cmd, ok := c.commands[state.Command.String()]; ok {
					action = cmd.Init
				}
			} else if command := c.commands[state.Command.String()]; command != nil && state.Action != Empty && command.Actions() != nil {
				if routes := command.Actions(); slices.Contains(update.Type(), routes[state.Action].AllowUpdates...) {
					action = routes[state.Action].Action
				}
			}
		}
		for _, middleware := range c.middlewares {
			action = middleware(action)
		}
		{
			next, err := action(ctx, update)
			if e := errors.Cast(err); e == nil {
				state.Action = next
			} else if !e.IsHandled() {
				return err
			}
		}
		if err := c.SetState(ctx, update.Chat().ID, state); err != nil {
			return err
		}
	}
	return nil
}
