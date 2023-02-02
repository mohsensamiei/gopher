package framework

import (
	"os"

	"github.com/fatih/color"
	"github.com/pinosell/gopher/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) up(cmd *cobra.Command, args []string) error {
	if err := os.Setenv("ENV_FILE", "../configs/.env"); err != nil {
		return err
	}

	color.Green("press ctrl+c to stop")
	if err := execext.CommandContextStream(cmd.Context(), "docker", "compose", "-f", "./deploy/docker-compose.deploy.yml", "up"); err != nil {
		return nil
	}
	return nil
}
