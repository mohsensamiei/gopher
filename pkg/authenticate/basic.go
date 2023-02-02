package authenticate

type Basic struct {
	Username string
	Password string
}

func NewBasic(username, password string) *Basic {
	return &Basic{
		Username: username,
		Password: password,
	}
}

func (t Basic) Type() Type {
	return BasicType
}
