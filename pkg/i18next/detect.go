package i18next

import (
	"github.com/gorilla/mux"
	"github.com/pinosell/gopher/pkg/httpext"
	"golang.org/x/text/language"
	"net/http"
	"strings"
)

func DetectMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			accept := httpext.Header(req, httpext.AcceptLanguageHeader)
			if accept == "" {
				next.ServeHTTP(res, req)
				return
			}
			lang, err := language.Parse(strings.Split(strings.Split(accept, ",")[0], ";")[0])
			if err != nil {
				next.ServeHTTP(res, req)
				return
			}
			next.ServeHTTP(res, req.WithContext(SetLang(req.Context(), lang)))
		})
	}
}
