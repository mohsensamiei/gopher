package framework

import (
	"time"

	"github.com/fatih/color"
	"github.com/mohsensamiei/gopher/v2/internal/pkg/helpers"
	"github.com/mohsensamiei/gopher/v2/pkg/cobraext"
	"github.com/spf13/cobra"
)

func NewCommander() cobraext.CommanderRegister {
	return &Commander{}
}

type Commander struct {
}

func (c Commander) RegisterCommander(root *cobra.Command) {
	{
		dump := &cobra.Command{
			Use:   "fmt",
			Args:  cobra.MaximumNArgs(1),
			Short: "Format project codes",
			RunE:  runE(c.fmt),
		}
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "dep",
			Args:  cobra.MaximumNArgs(1),
			Short: "Download project dependencies",
			RunE:  runE(c.dep),
		}
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "doc",
			Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
			Short: "Generate swagger document by general info file",
			RunE:  runE(c.doc),
		}
		cobraext.AddFlag(dump, "main", "m", "", "path of doc service main file, like: cmd/doc/main.go", true)
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "proto",
			Args:  cobra.MaximumNArgs(1),
			Short: "Generate golang codes from protobuf",
			RunE:  runE(c.proto),
		}
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "init",
			Args:  cobra.MaximumNArgs(1),
			Short: "Initial project repository codebase",
			RunE:  runE(c.init),
		}
		cobraext.AddFlag(dump, "rep", "", "", "project repository path, like: github.com/my/repo", true)
		cobraext.AddFlag(dump, "reg", "", "", "project registry path, like: ghcr.io/my/repo", true)
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "srv",
			Args:  cobra.MaximumNArgs(1),
			Short: "Add a new service to project",
			RunE:  runE(c.srv),
		}
		cobraext.AddFlag(dump, "name", "n", "", "name of service, like: container", true)
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "cmd",
			Args:  cobra.MaximumNArgs(1),
			Short: "Add a new command to service",
			RunE:  runE(c.cmd),
		}
		cobraext.AddFlag(dump, "srv", "s", "", "name of service, like: container", true)
		cobraext.AddFlag(dump, "name", "n", "", "name of command, like: finance", true)
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "app",
			Args:  cobra.MaximumNArgs(1),
			Short: "Add a new application to command",
			RunE:  runE(c.app),
		}
		cobraext.AddFlag(dump, "cmd", "c", "", "name of command, like: finance", true)
		cobraext.AddFlag(dump, "name", "n", "", "name of application, like: invoices", true)
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "migrate",
			Args:  cobra.MaximumNArgs(1),
			Short: "Add a new migration script to command",
			RunE:  runE(c.migrate),
		}
		cobraext.AddFlag(dump, "cmd", "c", "", "name of command, like: finance", true)
		cobraext.AddFlag(dump, "name", "n", "", "name of migration, like: create uuid extension", true)
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "lang",
			Args:  cobra.MaximumNArgs(1),
			Short: "Add a new language toml to project",
			RunE:  runE(c.lang),
		}
		cobraext.AddFlag(dump, "name", "n", "", "abbreviation of language, like: en", true)
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "build",
			Args:  cobra.MaximumNArgs(1),
			Short: "Build services docker image",
			RunE:  runE(c.build),
		}
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "test",
			Args:  cobra.MaximumNArgs(1),
			Short: "Tests project codes",
			RunE:  runE(c.test),
		}
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "run",
			Args:  cobra.MaximumNArgs(1),
			Short: "Build and run project services",
			RunE:  runE(c.run),
		}
		cobraext.AddFlag(dump, "cmd", "c", "", "name of command, like: finance", true)
		cobraext.AddFlag(dump, "env", "e", "configs/.env", "env file path", false)
		root.AddCommand(dump)
	}
	{
		dump := &cobra.Command{
			Use:   "up",
			Args:  cobra.MaximumNArgs(1),
			Short: "Start project services",
			RunE:  runE(c.up),
		}
		cobraext.AddFlag(dump, "env", "e", "configs/.env", "env file path", false)
		root.AddCommand(dump)
	}
}

func runE(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		if err := helpers.ChangeDirectory(args); err != nil {
			return err
		}
		if err := f(cmd, args); err != nil {
			return err
		}
		color.Green("Success in %v", time.Since(start))
		return nil
	}
}
