package strcaseext

import (
	"strings"
)

func ToDelimited(s string, sep string, f func(s string) string) string {
	return strings.Join(ToDelimitedSlice(s, sep, f), sep)
}

func ToDelimitedSlice(s string, sep string, f func(s string) string) []string {
	var res []string
	for _, v := range strings.Split(s, sep) {
		res = append(res, f(v))
	}
	return res
}
