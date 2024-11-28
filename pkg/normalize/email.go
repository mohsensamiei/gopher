package normalize

import (
	"github.com/mohsensamiei/gopher/v2/pkg/validatorext"
	"strings"
)

func Email(val string) (string, error) {
	val = strings.TrimSpace(val)
	if err := validatorext.IsEmail(val); err != nil {
		return "", err
	}
	return strings.ToLower(val), nil
}
