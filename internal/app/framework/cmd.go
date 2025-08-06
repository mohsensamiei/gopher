package framework

import (
	"fmt"
	"github.com/mohsensamiei/gopher/v3/internal/pkg/helpers"
	"github.com/mohsensamiei/gopher/v3/internal/pkg/templates"
	"github.com/mohsensamiei/gopher/v3/pkg/cobraext"
	"github.com/mohsensamiei/gopher/v3/pkg/slices"
	"github.com/spf13/cobra"
)

func (c Commander) cmd(cmd *cobra.Command, args []string) error {
	var command string
	if err := cobraext.Flag(cmd, "name", &command); err != nil {
		return err
	}
	if commands, err := helpers.Commands(); err != nil {
		return err
	} else if slices.Contains(command, commands...) {
		return fmt.Errorf("this command is already exists")
	}

	if err := helpers.MakeContents(map[string]string{
		fmt.Sprintf("cmd/%v/main.go", command): templates.CmdMain,
		fmt.Sprintf("configs/%v.env", command): templates.ConfigEnv,
	}); err != nil {
		return err
	}

	if err := c.dep(cmd, args); err != nil {
		return err
	}
	if err := c.fmt(cmd, args); err != nil {
		return err
	}
	return nil
}
