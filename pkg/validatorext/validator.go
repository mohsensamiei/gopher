package validatorext

import (
	"github.com/go-playground/validator/v10"
)

func New() *Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("decimal", Decimal)
	_ = validate.RegisterValidation("enum", Enum)
	_ = validate.RegisterValidation("phone", Phone)
	_ = validate.RegisterValidation("age", Age)
	return &Validate{
		Validate: validate,
	}
}

type Validate struct {
	*validator.Validate
}

func (v *Validate) SetDefault() {
	defaultValidate = v
}
