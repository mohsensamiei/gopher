package framework

import (
	"fmt"
	"os"

	"github.com/mohsensamiei/gopher/v2/internal/pkg/helpers"
	"github.com/mohsensamiei/gopher/v2/internal/pkg/templates"
	"github.com/spf13/cobra"
)

func (c Commander) lang(cmd *cobra.Command, _ []string) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	path := fmt.Sprintf("assets/locales/active.%v.toml", name)
	if _, err = os.Stat(path); err == nil {
		return fmt.Errorf("this language already exists")
	}

	if err = helpers.MakeContents(map[string]string{
		path: templates.LanguageToml,
	}, map[string]any{}); err != nil {
		return err
	}
	return nil
}
