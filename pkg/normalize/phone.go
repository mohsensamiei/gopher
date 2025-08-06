package normalize

import (
	"fmt"
	"github.com/mohsensamiei/gopher/v3/pkg/phonenumberext"
	"github.com/mohsensamiei/gopher/v3/pkg/validatorext"
)

func Phone(val string) (string, error) {
	if err := validatorext.IsPhone(val); err != nil {
		return "", err
	}
	phone, err := phonenumberext.Normalize(val)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("+%v", phone), nil
}
