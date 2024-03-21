package stringsext

import (
	"reflect"
	"sort"
	"strings"
)

func EqualSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	aCopy := make([]string, len(a))
	bCopy := make([]string, len(b))

	copy(aCopy, a)
	copy(bCopy, b)

	sort.Strings(aCopy)
	sort.Strings(bCopy)

	return reflect.DeepEqual(aCopy, bCopy)
}

func TrimSlice(s []string) []string {
	var r []string
	for _, v := range s {
		if tv := strings.TrimSpace(v); tv != "" {
			r = append(r, v)
		}
	}
	return r
}
