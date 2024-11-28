package validatorext

var (
	defaultValidate = New()
)

func Struct(s any, fields ...string) error {
	return defaultValidate.Struct(s, fields...)
}
