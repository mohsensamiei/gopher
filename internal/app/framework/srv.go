package framework

import (
	"fmt"
	"os"
	"strings"

	"github.com/pinosell/gopher/internal/pkg/helpers"
	"github.com/pinosell/gopher/internal/pkg/templates"
	"github.com/pinosell/gopher/pkg/cobraext"
	"github.com/pinosell/gopher/pkg/slices"
	"github.com/pinosell/gopher/pkg/templateext"
	"github.com/spf13/cobra"
)

func (c Commander) srv(cmd *cobra.Command, args []string) error {
	var (
		service  string
		services []string
	)
	services, err := helpers.Services()
	if err != nil {
		return err
	}
	if err = cobraext.Flag(cmd, "name", &service); err != nil {
		return err
	}
	if slices.Contains(service, services...) {
		return fmt.Errorf("this service is already exists")
	}

	var registry string
	registry, err = helpers.Registry("deploy/docker-compose.build.yml")
	if err != nil {
		return err
	}

	if err = helpers.MakeStructure([]string{
		fmt.Sprintf("services/%v", service),
	}); err != nil {
		return err
	}

	var deploy string
	deploy, err = appendServiceDeploy(registry, service)
	if err != nil {
		return err
	}

	if err = helpers.MakeContents(map[string]string{
		"deploy/docker-compose.build.yml":              deploy,
		fmt.Sprintf("services/%v/Dockerfile", service): templates.ServiceDockerfile,
	}, map[string]any{
		"name":    service,
		"service": "{{ .service }}",
		"command": "{{ .command }}",
	}); err != nil {
		return err
	}
	return nil
}

func appendServiceDeploy(registry, service string) (string, error) {
	serviceDeploy, err := templateext.Format(templates.DeployBuildService, map[string]any{
		"name":     service,
		"registry": registry,
		"service":  "{{ .service }}",
	})
	if err != nil {
		return "", err
	}

	var deployFile []byte
	deployFile, err = os.ReadFile("deploy/docker-compose.build.yml")
	if err != nil {
		return "", err
	}

	deployText := strings.ReplaceAll(string(deployFile), "# {{ .service }}", "{{ .service }}")
	deployText, err = templateext.Format(deployText, map[string]any{
		"service": strings.TrimSpace(serviceDeploy),
	})
	if err != nil {
		return "", err
	}
	return deployText, nil
}
