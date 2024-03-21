package query

var (
	Empty Encode = ""
)

type Encode string

func (q Query) Encode() Encode {
	return Encode(q.String())
}

func (e Encode) Parse() (*Query, error) {
	return Parse(string(e))
}
