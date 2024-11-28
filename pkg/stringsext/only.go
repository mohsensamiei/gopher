package stringsext

func Only(str string, length int) string {
	if runes := []rune(str); len(runes) > length {
		return string(runes[:length])
	}
	return str
}
