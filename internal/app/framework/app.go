package framework

import (
	"fmt"
	"os"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/mohsensamiei/gopher/internal/pkg/helpers"
	"github.com/mohsensamiei/gopher/internal/pkg/templates"
	"github.com/mohsensamiei/gopher/pkg/slices"
	"github.com/mohsensamiei/gopher/pkg/templateext"
	"github.com/spf13/cobra"
)

var (
	pluralSingular = pluralize.NewClient()
)

func (c Commander) app(cmd *cobra.Command, args []string) error {
	repository, err := helpers.Repository()
	if err != nil {
		return err
	}

	var (
		command  string
		commands []string
	)
	commands, err = helpers.Commands()
	if err != nil {
		return err
	}
	command, err = cmd.Flags().GetString("cmd")
	if err != nil {
		return err
	}
	if !slices.Contains(command, commands...) {
		return fmt.Errorf("command %v not found", command)
	}

	var (
		application  string
		applications []string
	)
	applications, err = helpers.Applications()
	if err != nil {
		return err
	}
	application, err = cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	if slices.Contains(application, applications...) {
		return fmt.Errorf("this application is already exists")
	}

	var (
		plural   string
		singular string
	)
	if pluralSingular.IsSingular(application) {
		singular = strcase.ToSnake(application)
		plural = strcase.ToSnake(pluralSingular.Plural(application))
	} else {
		plural = strcase.ToSnake(application)
		singular = strcase.ToSnake(pluralSingular.Singular(application))
	}

	if err = helpers.MakeStructure([]string{
		fmt.Sprintf("internal/app/%v", plural),
	}); err != nil {
		return err
	}

	var main string
	main, err = appendAppCommand(repository, command, plural)
	if err != nil {
		return err
	}

	if err = helpers.MakeContents(map[string]string{
		fmt.Sprintf("api/src/%v_model.proto", singular):      templates.ApiModel,
		fmt.Sprintf("api/src/%v_service.proto", singular):    templates.ApiService,
		fmt.Sprintf("internal/app/%v/controller.go", plural): templates.AppController,
		fmt.Sprintf("internal/app/%v/service.go", plural):    templates.AppService,
		fmt.Sprintf("cmd/%v/main.go", command):               main,
	}, map[string]any{
		"repository": repository,
		"Singular":   strcase.ToCamel(singular),
		"Plural":     strcase.ToCamel(plural),
		"singular":   singular,
		"plural":     plural,
		"command":    command,
		"import":     "{{ .import }}",
		"controller": "{{ .controller }}",
		"service":    "{{ .service }}",
	}); err != nil {
		return err
	}

	if err = c.proto(cmd, args); err != nil {
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

func appendAppCommand(repository, command, plural string) (string, error) {
	cmdImport, err := templateext.Format(templates.CmdImport, map[string]any{
		"plural":     plural,
		"repository": repository,
		"import":     "{{ .import }}",
	})
	if err != nil {
		return "", err
	}

	var cmdController string
	cmdController, err = templateext.Format(templates.CmdController, map[string]any{
		"plural":     plural,
		"repository": repository,
		"controller": "{{ .controller }}",
	})
	if err != nil {
		return "", err
	}

	var cmdService string
	cmdService, err = templateext.Format(templates.CmdService, map[string]any{
		"plural":     plural,
		"repository": repository,
		"service":    "{{ .service }}",
	})
	if err != nil {
		return "", err
	}

	var cmdFile []byte
	cmdFile, err = os.ReadFile(fmt.Sprintf("cmd/%v/main.go", command))
	if err != nil {
		return "", err
	}

	cmdText := string(cmdFile)
	cmdText = strings.ReplaceAll(cmdText, "// {{ .import }}", "{{ .import }}")
	cmdText = strings.ReplaceAll(cmdText, "// {{ .controller }}", "{{ .controller }}")
	cmdText = strings.ReplaceAll(cmdText, "// {{ .service }}", "{{ .service }}")
	cmdText, err = templateext.Format(cmdText, map[string]any{
		"import":     strings.TrimSpace(cmdImport),
		"controller": strings.TrimSpace(cmdController),
		"service":    strings.TrimSpace(cmdService),
	})
	if err != nil {
		return "", err
	}
	return cmdText, nil
}
