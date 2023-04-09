package validatorext

var (
	defaultValidate = New()
)

func Struct(s any) error {
	return defaultValidate.Struct(s)
}
