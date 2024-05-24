package framework

import (
	"github.com/mohsensamiei/gopher/v2/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) build(cmd *cobra.Command, _ []string) error {
	if err := execext.CommandContextStream(cmd.Context(), "docker", "compose", "-f", "deploy/docker-compose.build.yml", "build", "--no-cache"); err != nil {
		return err
	}
	return nil
}
