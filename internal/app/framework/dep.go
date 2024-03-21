package framework

import (
	"github.com/mohsensamiei/gopher/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) dep(cmd *cobra.Command, args []string) error {
	if err := execext.CommandContextStream(cmd.Context(), "go", "mod", "tidy"); err != nil {
		return err
	}
	if err := execext.CommandContextStream(cmd.Context(), "go", "mod", "vendor"); err != nil {
		return err
	}
	return nil
}
