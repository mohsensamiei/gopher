package framework

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mohsensamiei/gopher/v3/pkg/execext"
	"github.com/spf13/cobra"
)

func (c Commander) fmt(cmd *cobra.Command, _ []string) error {
	var shadow string
	if res, err := exec.CommandContext(cmd.Context(), "which", "shadow").CombinedOutput(); err == nil {
		shadow = strings.TrimSpace(string(res))
	}

	if shadow == "" {
		if err := execext.CommandContextStream(cmd.Context(), "go", "vet", "./..."); err != nil {
			return err
		}
	} else {
		if err := execext.CommandContextStream(cmd.Context(), "go", "vet", fmt.Sprintf("-vettool=%v", shadow), "./..."); err != nil {
			return err
		}
	}

	if err := execext.CommandContextStream(cmd.Context(), "go", "fmt", "./..."); err != nil {
		return err
	}
	return nil
}
