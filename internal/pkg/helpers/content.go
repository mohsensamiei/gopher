package helpers

import (
	"github.com/mohsensamiei/gopher/pkg/templateext"
	"os"
	"strings"
)

func MakeContents(contents map[string]string, data map[string]any) error {
	for path, format := range contents {
		text, err := templateext.Format(format, data)
		if err != nil {
			return err
		}
		if err = os.WriteFile(path, []byte(strings.TrimSpace(text)), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
