package strcaseext

import (
	"strings"
)

func Delimited(s string, sep string, f func(s string) string) string {
	return strings.Join(DelimitedSlice(s, sep, f), sep)
}

func DelimitedSlice(s string, sep string, f func(s string) string) []string {
	var res []string
	for _, v := range strings.Split(s, sep) {
		res = append(res, f(v))
	}
	return res
}
