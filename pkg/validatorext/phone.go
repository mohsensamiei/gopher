package validatorext

import (
	"github.com/go-playground/validator/v10"
	"github.com/pinosell/gopher/pkg/phonenumberext"
	"reflect"
)

func phone(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	v := fl.Field().String()
	_, err := phonenumberext.Normalize(v)
	return err == nil
}
