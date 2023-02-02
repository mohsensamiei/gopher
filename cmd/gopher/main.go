package main

import (
	"github.com/pinosell/gopher/internal/app/framework"
	"github.com/pinosell/gopher/pkg/cobraext"
	"github.com/pinosell/gopher/pkg/logext"
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
		Long:    "GOPHER\nSimple Golang Framework",
		Run:     cobraext.HelpFunc,
	}, []cobraext.CommanderRegister{
		framework.NewCommander(),
	})
}
