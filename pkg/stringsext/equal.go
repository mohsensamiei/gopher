package stringsext

import "strings"

func Comparable(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

func Equal(str1, str2 string) bool {
	return Comparable(str1) == Comparable(str2)
}

func Contains(str, sub string) bool {
	return strings.Contains(Comparable(str), Comparable(sub))
}

func SliceContains(str string, subs ...string) bool {
	for _, sub := range subs {
		if Contains(str, sub) {
			return true
		}
	}
	return false
}
