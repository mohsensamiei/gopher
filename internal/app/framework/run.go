package framework

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mohsensamiei/gopher/v3/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) run(cmd *cobra.Command, _ []string) error {
	command, err := cmd.Flags().GetString("cmd")
	if err != nil {
		return err
	}

	if err = godotenv.Load(fmt.Sprintf("configs/%v.env", command)); err != nil {
		return err
	}
	if err = execext.CommandContextStream(cmd.Context(), "go", "run", fmt.Sprintf("cmd/%v/main.go", command)); err != nil {
		return err
	}
	return nil
}
