package framework

import (
	"fmt"

	"github.com/pinosell/gopher/internal/pkg/helpers"
	"github.com/pinosell/gopher/internal/pkg/templates"
	"github.com/pinosell/gopher/pkg/cobraext"
	"github.com/pinosell/gopher/pkg/errors"
	"github.com/pinosell/gopher/pkg/execext"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
)

func (c Commander) init(cmd *cobra.Command, args []string) error {
	switch _, err := helpers.Repository(); errors.Code(err) {
	case codes.NotFound:
	case codes.OK:
		return fmt.Errorf("repository already init")
	default:
		return err
	}

	var repository string
	if err := cobraext.Flag(cmd, "rep", &repository); err != nil {
		return err
	}
	var registry string
	if err := cobraext.Flag(cmd, "reg", &registry); err != nil {
		return err
	}

	if err := helpers.MakeStructure([]string{
		"api/src",
		"assets/locales",
		"assets/migrations",
		"assets/statics",
		"assets/templates",
		"internal/app",
		"internal/pkg",
		"cmd",
		"pkg",
		"deploy",
		"scripts",
		"services/gateway",
		"configs",
		"docs",
		"tests",
	}); err != nil {
		return err
	}

	if err := helpers.MakeContents(map[string]string{
		"api/src/misc.proto":               templates.ApiMisc,
		"assets/locales/.gitkeep":          templates.GitKeep,
		"assets/migrations/.gitkeep":       templates.GitKeep,
		"assets/statics/.gitkeep":          templates.GitKeep,
		"assets/templates/.gitkeep":        templates.GitKeep,
		"deploy/docker-compose.build.yml":  templates.DeployBuild,
		"deploy/docker-compose.deploy.yml": templates.DeployUp,
		"internal/app/.gitkeep":            templates.GitKeep,
		"internal/pkg/.gitkeep":            templates.GitKeep,
		"configs/.gitkeep":                 templates.GitKeep,
		"configs/.env":                     templates.ConfigEnv,
		"docs/.gitkeep":                    templates.GitKeep,
		"tests/.gitkeep":                   templates.GitKeep,
		"cmd/.gitkeep":                     templates.GitKeep,
		"pkg/.gitkeep":                     templates.GitKeep,
		"scripts/.gitkeep":                 templates.GitKeep,
		"services/gateway/default.conf":    templates.GatewayConf,
		"services/gateway/Dockerfile":      templates.GatewayDockerfile,
		"Makefile":                         templates.RootMakefile,
		".gitlab-ci.yml":                   templates.GitCI,
		".gitignore":                       templates.GitIgnore,
	}, map[string]any{
		"registry":   registry,
		"repository": repository,
		"service":    "{{ .service }}",
		"command":    "{{ .command }}",
	}); err != nil {
		return err
	}

	if err := execext.CommandContextStream(cmd.Context(), "go", "mod", "init", repository); err != nil {
		return err
	}
	if err := c.proto(cmd, args); err != nil {
		return err
	}
	return nil
}
