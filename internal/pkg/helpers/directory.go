package helpers

import (
	"os"
)

func ChangeDirectory(args []string) error {
	if len(args) == 0 {
		return nil
	}
	if err := os.Chdir(args[0]); err != nil {
		return err
	}
	return nil
}

func MakeStructure(structure []string) error {
	for _, path := range structure {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
