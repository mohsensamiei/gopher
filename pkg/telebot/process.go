package telebot

import (
	"context"
	"fmt"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/mohsensamiei/gopher/v2/pkg/i18next"
	"github.com/mohsensamiei/gopher/v2/pkg/mapext"
	"github.com/mohsensamiei/gopher/v2/pkg/redisext"
	"github.com/mohsensamiei/gopher/v2/pkg/slices"
	"github.com/mohsensamiei/gopher/v2/pkg/telegram"
	"google.golang.org/grpc/codes"
	"strings"
	"time"
)

func parseCommand(str string) (string, []string) {
	dump := strings.Split(str, " ")
	command := strings.ToLower(strings.TrimSpace(dump[0]))
	if len(dump) > 1 {
		return command, dump[1:]
	}
	return command, nil
}

type State struct {
	Command string `json:"command,omitempty"`
	Action  string `json:"action,omitempty"`
	Data    []byte `json:"data,omitempty"`
}

func (c Client) recoveredProcess(ctx context.Context, update telegram.Update) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%s", rec)
		}
	}()
	return c.process(ctx, update)
}

func (c Client) process(ctx context.Context, update telegram.Update) error {
	chatID := update.ChatID()
	if chatID == 0 {
		return errors.New(codes.Unimplemented).WithDetailF("unsupported update type '%v'", update.Type())
	}
	{
		mutex := redisext.FromContext(ctx).NewMutex("telegram:locks", fmt.Sprint(chatID), time.Minute)
		if err := mutex.LockContext(ctx); err != nil {
			return err
		}
		defer func() {
			_, _ = mutex.UnlockContext(ctx)
		}()
	}
	var state = new(State)
	switch err := redisext.FromContext(ctx).Get(ctx, "telegram:stats", fmt.Sprint(chatID), state); errors.Code(err) {
	case codes.OK, codes.NotFound:
	case codes.InvalidArgument:
		_ = redisext.FromContext(ctx).Del(ctx, "telegram:stats", fmt.Sprint(chatID))
	default:
		return err
	}
	var action Action
	{
		var err error
		if update.IsCommand() || update.IsSimilarCommand(mapext.Keys(c.commands)) {
			state.Command, _ = parseCommand(update.Message.Text)
			if cmd, ok := c.commands[state.Command]; ok {
				action = cmd.Init
			} else {
				err = errors.New(codes.InvalidArgument)
			}
		} else if routes := c.commands[state.Command].Actions(); routes == nil || state.Action == "" {
			err = errors.New(codes.Aborted)
		} else {
			if slices.Contains(update.Type(), routes[state.Action].AllowUpdates...) {
				action = routes[state.Action].Action
			} else {
				err = errors.New(codes.InvalidArgument)
			}
		}
		if err != nil {
			return c.handleError(ctx, err, chatID, update.MessageID())
		}
	}
	for _, middleware := range c.Middlewares {
		if ok, err := middleware(ctx, update); err != nil {
			return c.handleError(ctx, err, chatID, update.MessageID())
		} else if !ok {
			return nil
		}
	}
	{
		var err error
		state.Action, state.Data, err = action(ctx, update, state.Data)
		if err != nil {
			return c.handleError(ctx, err, chatID, update.MessageID())
		}
	}
	if err := redisext.FromContext(ctx).Set(ctx, "telegram:stats", fmt.Sprint(chatID), state, 0); err != nil {
		return err
	}
	return nil
}

func (c Client) handleError(ctx context.Context, err error, chatID, MessageID int) error {
	e := errors.Cast(err)
	e = e.SetLocalize(i18next.ByContext(ctx, e.Slug()))
	_, _ = c.SendMessage(telegram.SendMessage{
		ChatID:           chatID,
		Text:             e.Localize(),
		ReplyToMessageID: MessageID,
	})
	if !e.IsHandled() {
		return err
	}
	return nil
}
