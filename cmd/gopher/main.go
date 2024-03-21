package main

import (
	"github.com/mohsensamiei/gopher/v2/internal/app/framework"
	"github.com/mohsensamiei/gopher/v2/pkg/cobraext"
	"github.com/mohsensamiei/gopher/v2/pkg/logext"
	"github.com/spf13/cobra"
)

const (
	Service = "gopher"
)

var (
	Version = "latest"
)

func init() {
	logext.Initial(Service, Version)
}

func main() {
	cobraext.Execute(&cobra.Command{
		Use:     Service,
		Version: Version,
		Long:    "GOPHER\nAdvanced Golang Framework",
		Run:     cobraext.HelpFunc,
	}, []cobraext.CommanderRegister{
		framework.NewCommander(),
	})
}
