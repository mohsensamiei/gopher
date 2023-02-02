package query

import (
	"net/url"
)

type Query map[string][]string

func (q Query) get(key string) []string {
	return q[key]
}

func (q Query) set(key, value string) {
	q[key] = []string{value}
}

func (q Query) add(key, value string) {
	q[key] = append(q[key], value)
}

func Parse(str string) (Query, error) {
	values, err := url.ParseQuery(str)
	if err != nil {
		return nil, err
	}
	return Query(values), nil
}

func (q Query) Encode() Encode {
	return Encode(url.Values(q).Encode())
}

func (q Query) String() string {
	return url.Values(q).Encode()
}
