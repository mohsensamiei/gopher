package validatorext

import (
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"google.golang.org/grpc/codes"
)

func (v *Validate) Struct(s any, fields ...string) error {
	var err error
	if len(fields) == 0 {
		err = v.Validate.Struct(s)
	} else {
		err = v.Validate.StructPartial(s, fields...)
	}
	if err != nil {
		var validations []*errors.Validation
		for _, item := range err.(validator.ValidationErrors) {
			validations = append(validations, &errors.Validation{
				Tag:   strcase.ToSnake(item.Tag()),
				Field: strcase.ToSnake(item.Field()),
				Param: item.Param(),
			})
		}
		return errors.Wrap(err, codes.InvalidArgument).
			WithValidations(validations)
	}
	return nil
}
