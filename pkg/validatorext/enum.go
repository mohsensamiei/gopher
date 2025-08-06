package validatorext

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

var (
	refEnumValidation = reflect.TypeOf((*EnumValidation)(nil)).Elem()
)

func Enum(fl validator.FieldLevel) bool {
	if !fl.Field().Type().Implements(refEnumValidation) {
		return false
	}
	v := fl.Field().Interface()
	return v.(EnumValidation).InRange(v)
}

type EnumValidation interface {
	InRange(v any) bool
}
