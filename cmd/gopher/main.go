package main

import (
	"github.com/mohsensamiei/gopher/v3/internal/app/framework"
	"github.com/mohsensamiei/gopher/v3/pkg/cobraext"
	"github.com/spf13/cobra"
)

func main() {
	cobraext.Execute(&cobra.Command{
		Use:     "gopher",
		Version: "v3",
		Long:    "GOPHER\nAdvanced Golang Framework",
		Run:     cobraext.HelpFunc,
	}, []cobraext.CommanderRegister{
		framework.NewCommander(),
	})
}
