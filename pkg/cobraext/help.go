package cobraext

import "github.com/spf13/cobra"

func HelpFunc(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, args)
}
