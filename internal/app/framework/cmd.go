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

func (c Commander) cmd(cmd *cobra.Command, args []string) error {
	registry, err := helpers.Registry("deploy/docker-compose.deploy.yml")
	if err != nil {
		return err
	}

	var (
		service  string
		services []string
	)
	services, err = helpers.Services()
	if err != nil {
		return err
	}
	if err = cobraext.Flag(cmd, "srv", &service); err != nil {
		return err
	}
	if !slices.Contains(service, services...) {
		return fmt.Errorf("service %v is not found", service)
	}

	var (
		command  string
		commands []string
	)
	commands, err = helpers.Commands()
	if err != nil {
		return err
	}
	if err = cobraext.Flag(cmd, "name", &command); err != nil {
		return err
	}
	if slices.Contains(command, commands...) {
		return fmt.Errorf("this command is already exists")
	}

	if err = helpers.MakeStructure([]string{
		fmt.Sprintf("assets/migrations/%v", command),
		fmt.Sprintf("cmd/%v", command),
	}); err != nil {
		return err
	}

	var docker string
	docker, err = appendDockerCommand(service, command)
	if err != nil {
		return err
	}

	var deploy string
	deploy, err = appendDeployCommand(registry, service, command)
	if err != nil {
		return err
	}

	var conf string
	conf, err = appendGatewayCommand(service, command)
	if err != nil {
		return err
	}

	if err = helpers.MakeContents(map[string]string{
		fmt.Sprintf("assets/migrations/%v/.gitkeep", command): templates.GitKeep,
		fmt.Sprintf("cmd/%v/main.go", command):                templates.CmdMain,
		fmt.Sprintf("services/%v/Dockerfile", service):        docker,
		"deploy/docker-compose.deploy.yml":                    deploy,
		"services/gateway/default.conf":                       conf,
	}, map[string]any{
		"name":       command,
		"command":    "{{ .command }}",
		"import":     "{{ .import }}",
		"controller": "{{ .controller }}",
		"service":    "{{ .service }}",
	}); err != nil {
		return err
	}

	if err = c.dep(cmd, args); err != nil {
		return err
	}
	if err = c.fmt(cmd, args); err != nil {
		return err
	}
	return nil
}

func appendGatewayCommand(service, command string) (string, error) {
	serviceBuild, err := templateext.Format(templates.GatewayCommand, map[string]any{
		"name":    command,
		"command": "{{ .command }}",
	})
	if err != nil {
		return "", err
	}

	var confFile []byte
	confFile, err = os.ReadFile("services/gateway/default.conf")
	if err != nil {
		if os.IsNotExist(err) {
			var text string
			text, err = templateext.Format(templates.GatewayConf, map[string]any{
				"command": "{{ .command }}",
			})
			if err != nil {
				return "", err
			}
			confFile = []byte(strings.TrimSpace(text))
		} else {
			return "", err
		}
	}
	confText := strings.ReplaceAll(string(confFile), "# {{ .command }}", "{{ .command }}")
	confText, err = templateext.Format(confText, map[string]any{
		"command": strings.TrimSpace(serviceBuild),
	})
	if err != nil {
		return "", err
	}
	return confText, nil
}

func appendDockerCommand(service, command string) (string, error) {
	serviceBuild, err := templateext.Format(templates.ServiceBuild, map[string]any{
		"name":    command,
		"command": "{{ .command }}",
	})
	if err != nil {
		return "", err
	}

	var dockerFile []byte
	dockerFile, err = os.ReadFile(fmt.Sprintf("services/%v/Dockerfile", service))
	if err != nil {
		return "", err
	}

	dockerText := strings.ReplaceAll(string(dockerFile), "# {{ .command }}", "{{ .command }}")
	dockerText, err = templateext.Format(dockerText, map[string]any{
		"command": strings.TrimSpace(serviceBuild),
	})
	if err != nil {
		return "", err
	}
	return dockerText, nil
}

func appendDeployCommand(registry, service, command string) (string, error) {
	serviceDeploy, err := templateext.Format(templates.DeployRunService, map[string]any{
		"name":     command,
		"service":  service,
		"registry": registry,
		"command":  "{{ .command }}",
	})
	if err != nil {
		return "", err
	}

	var deployFile []byte
	deployFile, err = os.ReadFile("deploy/docker-compose.deploy.yml")
	if err != nil {
		return "", err
	}

	deployText := strings.ReplaceAll(string(deployFile), "# {{ .command }}", "{{ .command }}")
	deployText, err = templateext.Format(deployText, map[string]any{
		"command": strings.TrimSpace(serviceDeploy),
	})
	if err != nil {
		return "", err
	}
	return deployText, nil
}
