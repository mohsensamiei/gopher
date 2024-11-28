package validatorext

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
)

func decimal(fl validator.FieldLevel) bool {
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
		var minimum int64
		minimum, err = strconv.ParseInt(fl.Param(), 10, 64)
		if err != nil {
			panic("invalid min argument")
		}
		if dec < minimum {
			return false
		}
	}
	return true
}
