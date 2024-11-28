package telebot

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/mohsensamiei/gopher/v2/pkg/i18next"
	"github.com/mohsensamiei/gopher/v2/pkg/telegram"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
)

type Middleware func(next Action) Action

func abortMiddleware(next Action) Action {
	return func(ctx context.Context, update telegram.Update) (Keyword, error) {
		if next == nil {
			if update.Chat().Type == telegram.Private {
				return Empty, errors.New(codes.Aborted)
			} else {
				return Empty, nil
			}
		}
		return next(ctx, update)
	}
}

func ActionMiddleware(action telegram.Action) Middleware {
	return func(next Action) Action {
		return func(ctx context.Context, update telegram.Update) (Keyword, error) {
			if chatID := update.Chat().ID; chatID > 0 {
				_, _ = telegram.FromContext(ctx).SendAction(telegram.SendAction{
					ChatID: chatID,
					Action: action,
				})
			}
			return next(ctx, update)
		}
	}
}

func LangMiddleware(next Action) Action {
	return func(ctx context.Context, update telegram.Update) (Keyword, error) {
		if update.From() == nil {
			return Empty, nil
		}
		var base language.Base
		if tags, _, err := language.ParseAcceptLanguage(update.From().LanguageCode); err != nil {
			return Empty, err
		} else {
			base, _ = tags[0].Base()
		}
		lang, err := language.Parse(base.String())
		if err != nil {
			return Empty, err
		}
		return next(i18next.SetLang(ctx, lang), update)
	}
}

func ErrorMiddleware(next Action) Action {
	return func(ctx context.Context, update telegram.Update) (Keyword, error) {
		keyword, err := next(ctx, update)
		if err != nil {
			e := errors.Cast(err)
			e = e.SetLocalize(i18next.ByContext(ctx, e.Slug()))
			_, _ = telegram.FromContext(ctx).SendMessage(telegram.SendMessage{
				ChatID:           update.Chat().ID,
				Text:             e.Localize(),
				ReplyToMessageID: update.MessageID(),
			})
		}
		return keyword, err
	}
}

func (c *Client) Middleware(middleware ...Middleware) {
	c.middlewares = append(c.middlewares, middleware...)
}
