package stringsext

import "strings"

const Empty = ""

func IsNilOrEmpty[T ~string | *string](input T) bool {
	switch v := any(input).(type) {
	case string:
		return strings.TrimSpace(v) == ""
	case *string:
		return v == nil || strings.TrimSpace(*v) == ""
	default:
		return false
	}
}
