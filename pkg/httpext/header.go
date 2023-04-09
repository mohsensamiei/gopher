package httpext

import "net/http"

func RealIP(req *http.Request) string {
	return Header(req, RealIPHeader)
}

func Header(req *http.Request, key string) string {
	return req.Header.Get(key)
}
