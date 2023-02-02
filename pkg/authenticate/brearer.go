package authenticate

type Bearer struct {
	Token string
}

func NewBearer(token string) *Bearer {
	return &Bearer{
		token,
	}
}

func (t Bearer) Type() Type {
	return BearerType
}
