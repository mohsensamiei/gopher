package authenticate

type Type string

const (
	BearerType Type = "Bearer"
	BasicType  Type = "Basic"
)

type Authenticate interface {
	Type() Type
}
