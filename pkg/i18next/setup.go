package i18next

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Setup(configs Configs, path string) error {
	defaultLang = configs.DefaultLang

	bundle := i18n.NewBundle(defaultLang)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, "toml") {
			return nil
		}
		if _, err = bundle.LoadMessageFile(path); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	for _, tag := range bundle.LanguageTags() {
		languages[tag] = i18n.NewLocalizer(bundle, tag.String())
	}
	return nil
}
