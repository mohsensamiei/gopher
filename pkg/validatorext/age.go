package validatorext

import (
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

func Age(fl validator.FieldLevel) bool {
	param, err := strconv.Atoi(fl.Param())
	if err != nil || param < 0 {
		panic("invalid age argument")
	}
	switch v := fl.Field().Interface().(type) {
	case time.Time:
		return calculateAge(v) >= param
	case *time.Time:
		return calculateAge(*v) >= param
	}
	return false
}

func calculateAge(birthDate time.Time) int {
	today := time.Now()
	age := today.Year() - birthDate.Year()
	if today.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}
