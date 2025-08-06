package helpers

import "os"

func ChangeDirectory(args []string) error {
	if len(args) == 0 {
		return nil
	}
	if err := os.Chdir(args[0]); err != nil {
		return err
	}
	return nil
}
