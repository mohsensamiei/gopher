package cobraext

import "github.com/spf13/cobra"

type CommanderRegister interface {
	RegisterCommander(root *cobra.Command)
}
