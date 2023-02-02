package i18next

import (
	"golang.org/x/text/language"
)

type Configs struct {
	DefaultLang language.Tag `env:"DEFAULT_LANG" envDefault:"en"`
}
