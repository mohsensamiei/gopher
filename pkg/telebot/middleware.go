package telebot

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/i18next"
	"github.com/mohsensamiei/gopher/v3/pkg/telegram"
	"golang.org/x/text/language"
)

type Middleware func(next Action) Action

func ActionMiddleware(action telegram.Action) Middleware {
	return func(next Action) Action {
		return func(ctx context.Context, update telegram.Update) (Keyword, error) {
			if chatID := update.Chat().ID; chatID > 0 {
				_, _ = di.Provide[*telegram.Connection](ctx).SendAction(telegram.SendAction{
					ChatID: chatID,
					Action: action,
				})
			}
			return next(ctx, update)
		}
	}
}

func LangMiddleware() Middleware {
	return func(next Action) Action {
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
}

func ErrorMiddleware() Middleware {
	return func(next Action) Action {
		return func(ctx context.Context, update telegram.Update) (Keyword, error) {
			keyword, err := next(ctx, update)
			if err != nil {
				e := errors.Cast(err)
				e = e.SetLocalize(i18next.ByContext(ctx, e.Slug()))
				_, _ = di.Provide[*telegram.Connection](ctx).SendMessage(telegram.SendMessage{
					ChatID:           update.Chat().ID,
					Text:             e.Localize(),
					ReplyToMessageID: update.MessageID(),
				})
			}
			return keyword, err
		}
	}
}

func (c *Client) Middleware(middleware ...Middleware) {
	c.middlewares = append(c.middlewares, middleware...)
}
