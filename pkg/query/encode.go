package query

type Encode string

const (
	Empty Encode = ""
)

func (s Encode) String() string {
	return string(s)
}

func (s Encode) Parse() (Query, error) {
	return Parse(string(s))
}
