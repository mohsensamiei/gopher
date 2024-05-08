package framework

import (
	"fmt"
	"github.com/mohsensamiei/gopher/v2/internal/pkg/templates"
	"github.com/mohsensamiei/gopher/v2/pkg/templateext"
	"os"
	"regexp"
	"strings"

	"github.com/mohsensamiei/gopher/v2/pkg/execext"
	"github.com/spf13/cobra"
)

var (
	enumRegex       = regexp.MustCompile("type (.*) int32")
	enumImportRegex = regexp.MustCompile("import \\(((.|\\n)*)\\)\\n\\nconst")
)

func (c Commander) proto(cmd *cobra.Command, args []string) error {
	if err := execext.CommandContextStream(cmd.Context(), "rm", "-f", "api/*.pb.go"); err != nil {
		return err
	}
	defer func() {
		_ = execext.CommandContextStream(cmd.Context(), "rm", "-f", "api/src/*.pb.go")
	}()

	files, err := os.ReadDir("api/src")
	if err != nil {
		return err
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".proto") {
			continue
		}
		if err = execext.CommandContextStream(cmd.Context(), "protoc", "--go_out=.", "--go_opt=paths=source_relative", "--go-grpc_opt=paths=source_relative", "--go-grpc_out=require_unimplemented_servers=false:.", fmt.Sprintf("api/src/%v", file.Name())); err != nil {
			return err
		}
	}

	files, err = os.ReadDir("api/src")
	if err != nil {
		return err
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".pb.go") {
			continue
		}
		filePath := fmt.Sprintf("api/src/%v", file.Name())
		{
			var bin []byte
			bin, err = os.ReadFile(filePath)
			if err != nil {
				return err
			}

			body := string(bin)
			matches := enumRegex.FindAllStringSubmatch(body, -1)
			if len(matches) > 0 {
				im := fmt.Sprintf("%v%v", templates.ApiEnumImport, enumImportRegex.FindStringSubmatch(body)[1])
				body = strings.ReplaceAll(body, enumImportRegex.FindStringSubmatch(body)[1], im)
			}
			for _, name := range matches {
				var add string
				add, err = templateext.Format(templates.ApiEnum, map[string]any{
					"Enum": name[1],
				})
				body = fmt.Sprintf("%v\n\n%v", body, add)
			}
			bin = []byte(body)

			if err = os.WriteFile(filePath, bin, os.ModePerm); err != nil {
				return err
			}
		}
		if err = os.Rename(filePath, fmt.Sprintf("api/%v", file.Name())); err != nil {
			return err
		}
		if err = execext.CommandContextStream(cmd.Context(), "protoc-go-inject-tag", fmt.Sprintf("-input=api/%v", file.Name()), "-XXX_skip=json,xml,yaml"); err != nil {
			return err
		}
	}
	return nil
}
