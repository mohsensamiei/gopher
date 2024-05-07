package framework

import (
	"github.com/mohsensamiei/gopher/v2/pkg/cobraext"
	"github.com/mohsensamiei/gopher/v2/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) doc(cmd *cobra.Command, args []string) error {
	var doc string
	if err := cobraext.Flag(cmd, "doc", &doc); err != nil {
		return err
	}
	if err := c.fmt(cmd, args); err != nil {
		return err
	}
	if err := execext.CommandContextStream(cmd.Context(), "swag", "init",
		"--generalInfo", doc,
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
