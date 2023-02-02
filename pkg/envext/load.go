package envext

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadIfExists(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if err := godotenv.Load(filename); err != nil {
		return err
	}
	return nil
}
