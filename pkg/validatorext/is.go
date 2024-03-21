package validatorext

type isEmail struct {
	Email string `validate:"max=320,email,required"`
}

func IsEmail(v string) error {
	return defaultValidate.Struct(&isEmail{
		Email: v,
	})
}

type isMobile struct {
	Mobile string `validate:"max=15,phone,required"`
}

func IsMobile(v string) error {
	return defaultValidate.Struct(&isMobile{
		Mobile: v,
	})
}
