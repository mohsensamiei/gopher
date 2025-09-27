package framework

import (
	"fmt"
	"log"

	"github.com/mohsensamiei/gopher/v3/internal/pkg/helpers"
	"github.com/mohsensamiei/gopher/v3/internal/pkg/templates"
	"github.com/mohsensamiei/gopher/v3/pkg/cobraext"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/execext"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
)

func (c Commander) init(cmd *cobra.Command, args []string) error {
	switch e, err := helpers.Repository(); errors.Code(err) {
	case codes.NotFound:
	case codes.OK:
		log.Print(e)
		return fmt.Errorf("repository already init")
	default:
		return err
	}

	var repository string
	if err := cobraext.Flag(cmd, "rep", &repository); err != nil {
		return err
	}

	if err := helpers.MakeContents(map[string]string{
		"api/src/.gitkeep":           templates.GitKeep,
		"assets/locales/.gitkeep":    templates.GitKeep,
		"assets/migrations/.gitkeep": templates.GitKeep,
		"assets/statics/.gitkeep":    templates.GitKeep,
		"assets/templates/.gitkeep":  templates.GitKeep,
		"deploy/Dockerfile":          templates.DeployDockerfile,
		"internal/app/.gitkeep":      templates.GitKeep,
		"internal/pkg/.gitkeep":      templates.GitKeep,
		"configs/.gitkeep":           templates.GitKeep,
		"docs/init.go":               templates.DocInit,
		"tests/.gitkeep":             templates.GitKeep,
		"cmd/.gitkeep":               templates.GitKeep,
		"pkg/.gitkeep":               templates.GitKeep,
		"scripts/.gitkeep":           templates.GitKeep,
		"Makefile":                   templates.RootMakefile,
		".air.toml":                  templates.RootAir,
		".gitignore":                 templates.GitIgnore,
		".dockerignore":              templates.DockerIgnore,
	}); err != nil {
		return err
	}

	if err := execext.CommandContextStream(cmd.Context(), "go", "mod", "init", repository); err != nil {
		return err
	}
	return nil
}
