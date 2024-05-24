package framework

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mohsensamiei/gopher/v2/pkg/execext"
	"github.com/spf13/cobra"
	"os"
)

func (c Commander) run(cmd *cobra.Command, _ []string) error {
	command, err := cmd.Flags().GetString("cmd")
	if err != nil {
		return err
	}

	var envFile string
	envFile, err = cmd.Flags().GetString("env")
	if err != nil {
		return err
	}
	if err = os.Setenv("ENV_FILE", envFile); err != nil {
		return err
	}
	if err = godotenv.Load(envFile); err != nil {
		return err
	}

	if err = execext.CommandContextStream(cmd.Context(), "go", "run", fmt.Sprintf("cmd/%v/main.go", command)); err != nil {
		return err
	}
	return nil
}
