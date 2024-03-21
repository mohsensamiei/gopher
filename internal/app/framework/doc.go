package framework

import (
	"github.com/mohsensamiei/gopher/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) doc(cmd *cobra.Command, args []string) error {
	if err := c.fmt(cmd, args); err != nil {
		return err
	}
	if err := execext.CommandContextStream(cmd.Context(), "swag", "init",
		"--generalInfo", args[len(args)-1],
		"--output", "docs",
		"--propertyStrategy", "snakecase",
		"--outputTypes", "go",
		"--parseInternal",
		"--parseVendor",
		"--quiet",
	); err != nil {
		return err
	}
	return nil
}
