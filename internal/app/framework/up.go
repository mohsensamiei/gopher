package framework

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/mohsensamiei/gopher/v2/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) up(cmd *cobra.Command, _ []string) error {
	envFile, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}
	if err = os.Setenv("ENV_FILE", filepath.Join("..", envFile)); err != nil {
		return err
	}

	color.Green("press ctrl+c to stop")
	if err = execext.CommandContextStream(cmd.Context(), "docker", "compose", "-f", "deploy/docker-compose.deploy.yml", "up"); err != nil {
		return nil
	}
	return nil
}
