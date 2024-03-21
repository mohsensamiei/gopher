package framework

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mohsensamiei/gopher/v2/internal/pkg/helpers"
	"github.com/mohsensamiei/gopher/v2/internal/pkg/templates"
	"github.com/mohsensamiei/gopher/v2/pkg/cobraext"
	"github.com/mohsensamiei/gopher/v2/pkg/slices"
	"github.com/spf13/cobra"
)

func (c Commander) migrate(cmd *cobra.Command, args []string) error {
	var command string
	commands, err := helpers.Commands()
	if err != nil {
		return err
	}
	if err = cobraext.Flag(cmd, "cmd", &command); err != nil {
		return err
	}
	if !slices.Contains(command, commands...) {
		return fmt.Errorf("command %v not found", command)
	}

	var name string
	if err = cobraext.Flag(cmd, "name", &name); err != nil {
		return err
	}

	var num int
	num, err = helpers.MigrationNumber(command)
	if err != nil {
		return err
	}

	if err = helpers.MakeContents(map[string]string{
		fmt.Sprintf("assets/migrations/%v/%v_%v.up.sql", command, num, strcase.ToSnake(strings.TrimSpace(name))):   templates.MigrationUp,
		fmt.Sprintf("assets/migrations/%v/%v_%v.down.sql", command, num, strcase.ToSnake(strings.TrimSpace(name))): templates.MigrationDown,
	}, map[string]any{}); err != nil {
		return err
	}
	return nil
}
