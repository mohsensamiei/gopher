package httpext

import "net/http"

func RealIP(req *http.Request) string {
	return req.Header.Get("X-Real-IP")
}
