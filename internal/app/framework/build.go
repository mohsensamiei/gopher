package framework

import (
	"github.com/mohsensamiei/gopher/pkg/cobraext"
	"github.com/mohsensamiei/gopher/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) build(cmd *cobra.Command, args []string) error {
	if err := c.dep(cmd, args); err != nil {
		return err
	}
	if err := c.fmt(cmd, args); err != nil {
		return err
	}

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
