package framework

import (
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/mohsensamiei/gopher/v3/internal/pkg/helpers"
	"github.com/mohsensamiei/gopher/v3/internal/pkg/templates"
	"github.com/mohsensamiei/gopher/v3/pkg/cobraext"
	"github.com/mohsensamiei/gopher/v3/pkg/slices"
	"github.com/spf13/cobra"
)

var (
	pluralSingular = pluralize.NewClient()
)

func (c Commander) app(cmd *cobra.Command, args []string) error {
	if _, err := helpers.Repository(); err != nil {
		return err
	}

	var application string
	if err := cobraext.Flag(cmd, "name", &application); err != nil {
		return err
	}
	if pluralSingular.IsSingular(application) {
		application = pluralSingular.Plural(application)
	}
	application = strcase.ToSnake(application)

	if applications, err := helpers.Applications(); err != nil {
		return err
	} else if slices.Contains(application, applications...) {
		return fmt.Errorf("this application is already exists")
	}

	if err := helpers.MakeContents(map[string]string{
		fmt.Sprintf("internal/app/%v/.gitkeep", application): templates.GitKeep,
	}); err != nil {
		return err
	}
	return nil
}
