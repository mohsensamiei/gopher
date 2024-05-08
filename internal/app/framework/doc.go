package framework

import (
	"github.com/mohsensamiei/gopher/v2/pkg/cobraext"
	"github.com/mohsensamiei/gopher/v2/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) doc(cmd *cobra.Command, args []string) error {
	var main string
	if err := cobraext.Flag(cmd, "main", &main); err != nil {
		return err
	}
	if err := execext.CommandContextStream(cmd.Context(), "swag", "fmt"); err != nil {
		return err
	}
	if err := execext.CommandContextStream(cmd.Context(), "swag", "init",
		"--generalInfo", main,
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
