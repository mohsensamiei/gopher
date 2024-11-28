package i18next

import (
	"context"
	"github.com/fatih/structs"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func ByLangWithData(lang language.Tag, id string, data any) string {
	if structs.IsStruct(data) {
		data = structs.Map(data)
	}
	local, ok := languages[lang]
	if !ok {
		if local, ok = languages[defaultLang]; !ok {
			return id
		}
	}
	str, err := local.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
	})
	if err != nil {
		return id
	}
	return str
}

func ByLang(lang language.Tag, id string) string {
	return ByLangWithData(lang, id, nil)
}

func ByContext(ctx context.Context, id string) string {
	return ByContextWithData(ctx, id, nil)
}

func ByContextWithData(ctx context.Context, id string, data any) string {
	return ByLangWithData(GetLang(ctx), id, data)
}
