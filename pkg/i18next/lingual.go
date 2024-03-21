package i18next

import (
	"context"
	"golang.org/x/text/language"
)

type Lingual interface {
	GetLanguage() language.Tag
}

func ToContext(ctx context.Context, l Lingual) context.Context {
	return SetLang(ctx, l.GetLanguage())
}
