package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func MakeContents(contents map[string]string) error {
	for path, body := range contents {
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(strings.TrimSpace(body)), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
