package validatorext

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pinosell/gopher/pkg/errors"
	"google.golang.org/grpc/codes"
	"reflect"
	"regexp"
	"strconv"
)

const (
	mobileRegexString     = `^(((\+|00)(90|963|964|971|98))|0)?([1-9]\d*)$`
	postalCodeRegexString = `^[1-9]\d{9}$`
)

var (
	validate          = validator.New()
	refEnumValidation = reflect.TypeOf((*EnumValidation)(nil)).Elem()
	mobileRegex       = regexp.MustCompile(mobileRegexString)
	postalCodeRegex   = regexp.MustCompile(postalCodeRegexString)
)

func init() {
	_ = validate.RegisterValidation("decimal", ValidateDecimal)
	_ = validate.RegisterValidation("enum", ValidateEnum)
	_ = validate.RegisterValidation("mobile", ValidateMobile)
	_ = validate.RegisterValidation("postal_code", ValidatePostalCode)
}

type EnumValidation interface {
	InRange(v any) bool
}

func ValidateDecimal(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	v := fl.Field().String()
	dec, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return false
	}
	if fmt.Sprint(dec) != v {
		return false
	}
	if fl.Param() != "" {
		var min int64
		min, err = strconv.ParseInt(fl.Param(), 10, 64)
		if err != nil {
			panic("invalid min argument")
		}
		if dec < min {
			return false
		}
	}
	return true
}

func ValidateEnum(fl validator.FieldLevel) bool {
	if !fl.Field().Type().Implements(refEnumValidation) {
		return false
	}
	v := fl.Field().Interface()
	return v.(EnumValidation).InRange(v)
}

func ValidateMobile(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	v := fl.Field().String()
	return mobileRegex.MatchString(v)
}

func ValidatePostalCode(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	v := fl.Field().String()
	return postalCodeRegex.MatchString(v)
}

func Struct(s any) error {
	if err := validate.Struct(s); err != nil {
		return errors.New(codes.InvalidArgument).
			WithDetails(err.Error())
	}
	return nil
}
