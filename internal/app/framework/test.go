package framework

import (
	"github.com/mohsensamiei/gopher/v3/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) test(cmd *cobra.Command, _ []string) error {
	if err := execext.CommandContextStream(cmd.Context(), "go", "test", "-v", "./..."); err != nil {
		return err
	}
	return nil
}
