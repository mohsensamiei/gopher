package i18next

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	defaultLang = language.Und
	languages   = make(map[language.Tag]*i18n.Localizer)
)

func Languages() (tags []language.Tag) {
	for tag := range languages {
		tags = append(tags, tag)
	}
	return
}

func DefaultLang() language.Tag {
	return defaultLang
}
