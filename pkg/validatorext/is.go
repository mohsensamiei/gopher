package validatorext

type isEmail struct {
	Email string `validate:"max=320,email,required"`
}

func IsEmail(v string) error {
	return defaultValidate.Struct(&isEmail{
		Email: v,
	})
}

type isPhone struct {
	Phone string `validate:"max=15,phone,required"`
}

func IsPhone(v string) error {
	return defaultValidate.Struct(&isPhone{
		Phone: v,
	})
}
