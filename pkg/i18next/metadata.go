package i18next

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/metadataext"
	"golang.org/x/text/language"
)

const (
	key = "lang"
)

func GetLang(ctx context.Context) language.Tag {
	if lang, ok := metadataext.GetValue(ctx, key); ok {
		tag, err := language.Parse(lang)
		if err != nil {
			return defaultLang
		}
		return tag
	}
	return defaultLang
}

func SetLang(ctx context.Context, lang language.Tag) context.Context {
	if lang == language.Und {
		return ctx
	}
	return metadataext.SetValue(ctx, key, lang.String())
}
