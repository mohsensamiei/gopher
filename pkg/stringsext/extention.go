package stringsext

import (
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func Generate(alphabet string, length int) string {
	bytes := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for index, cache, remain := length-1, src.Int63(), letterIdxMax; index >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(alphabet) {
			bytes[index] = alphabet[idx]
			index--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&bytes))
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
