package httpext

import (
	"github.com/mohsensamiei/gopher/v3/pkg/mimeext"
	"net/http"
	"strings"
)

func Header(req *http.Request, key string) string {
	return req.Header.Get(key)
}

func ParseHeader[T any](req *http.Request, key string, parse func(str string) (T, error)) (T, error) {
	return parse(Header(req, key))
}

func GetContentMIME(header http.Header) string {
	mime := header.Get(ContentTypeHeader)
	if mime == "application/octet-stream" {
		if pdf, _ := mimeext.ExtensionByType(mimeext.PDF); strings.Contains(strings.ToLower(header.Get(ContentDisposition)), pdf) {
			return mimeext.PDF
		}
	}
	return mime
}
