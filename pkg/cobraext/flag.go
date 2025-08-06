package cobraext

import (
	"github.com/mohsensamiei/gopher/v3/pkg/pointers"
	"github.com/spf13/cobra"
)

func Flag[T string](cmd *cobra.Command, name string, result *T) error {
	switch any(result).(type) {
	case *string:
		res, err := cmd.Flags().GetString(name)
		if err != nil {
			return err
		}
		*result = *pointers.Pointer(T(res))
	}
	return nil
}

func AddFlag[T string](cmd *cobra.Command, name, shorthand, value, description string, required bool) {
	cmd.PersistentFlags().StringP(name, shorthand, value, description)
	if required {
		_ = cmd.MarkPersistentFlagRequired(name)
	}
}
