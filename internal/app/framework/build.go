package framework

import (
	"github.com/mohsensamiei/gopher/v2/pkg/cobraext"
	"github.com/mohsensamiei/gopher/v2/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) build(cmd *cobra.Command, args []string) error {
	var service string
	_ = cobraext.Flag(cmd, "srv", &service)
	if service == "" {
		if err := execext.CommandContextStream(cmd.Context(), "docker", "compose", "-f", "./deploy/docker-compose.build.yml", "build"); err != nil {
			return err
		}
	} else {
		if err := execext.CommandContextStream(cmd.Context(), "docker", "compose", "-f", "./deploy/docker-compose.build.yml", "build", service); err != nil {
			return err
		}
	}
	return nil
}
