package validatorext

import "github.com/go-playground/validator/v10"

func enum(fl validator.FieldLevel) bool {
	if !fl.Field().Type().Implements(refEnumValidation) {
		return false
	}
	v := fl.Field().Interface()
	return v.(EnumValidation).InRange(v)
}
