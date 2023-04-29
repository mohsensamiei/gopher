package validatorext

import (
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/pinosell/gopher/pkg/errors"
	"google.golang.org/grpc/codes"
	"reflect"
	"regexp"
	"strings"
)

type EnumValidation interface {
	InRange(v any) bool
}

var (
	refEnumValidation = reflect.TypeOf((*EnumValidation)(nil)).Elem()
	raw               = regexp.MustCompile(`Key: '([^']*)' Error:Field validation for '([^']*)' failed on the '([^']*)' tag`)
)

func New() *Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("decimal", decimal)
	_ = validate.RegisterValidation("enum", enum)
	_ = validate.RegisterValidation("phone", phone)
	return &Validate{
		Validate: validate,
	}
}

type Validate struct {
	*validator.Validate
}

func (v *Validate) Struct(s any) error {
	if err := v.Validate.Struct(s); err != nil {
		var validations []*errors.Validation
		for _, row := range raw.FindAllStringSubmatch(err.Error(), -1) {
			validations = append(validations, &errors.Validation{
				Tag:   strings.Split(strcase.ToSnake(row[3]), "_")[0],
				Field: strcase.ToSnake(row[2]),
			})
		}
		return errors.Wrap(err, codes.InvalidArgument).
			WithValidations(validations)
	}
	return nil
}

func (v *Validate) SetDefault() {
	defaultValidate = v
}
