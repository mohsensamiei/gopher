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
	enumRegex = regexp.MustCompile("type (.*) int32")
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
		{
			var bin []byte
			bin, err = os.ReadFile(file.Name())
			if err != nil {
				return err
			}

			body := string(bin)
			for _, name := range enumRegex.FindStringSubmatch(body) {
				var add string
				add, err = templateext.Format(templates.ApiEnum, map[string]any{
					"Enum": name,
				})
				body = fmt.Sprintf("%v\n%v", body, add)
			}
			bin = []byte(body)

			if err = os.WriteFile(file.Name(), bin, os.ModePerm); err != nil {
				return err
			}
		}
		if err = os.Rename(fmt.Sprintf("api/src/%v", file.Name()), fmt.Sprintf("api/%v", file.Name())); err != nil {
			return err
		}
		if err = execext.CommandContextStream(cmd.Context(), "protoc-go-inject-tag", fmt.Sprintf("-input=api/%v", file.Name()), "-XXX_skip=json,xml,yaml"); err != nil {
			return err
		}
	}

	if err = c.dep(cmd, args); err != nil {
		return err
	}
	return nil
}
