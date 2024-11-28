package base32ext

import (
	"encoding/base32"
	"strings"
)

func EncodeString(decoded string) string {
	return strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(decoded)))
}

func DecodeString(encoded string) (string, error) {
	bin, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(strings.ToUpper(encoded))
	if err != nil {
		return "", err
	}
	return string(bin), nil
}
