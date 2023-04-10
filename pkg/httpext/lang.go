package httpext

import (
	"github.com/gorilla/mux"
	"github.com/pinosell/gopher/pkg/i18next"
	"golang.org/x/text/language"
	"net/http"
	"strings"
)

func DetectLangMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			accept := Header(req, AcceptLanguageHeader)
			if accept == "" {
				next.ServeHTTP(res, req)
				return
			}
			lang, err := language.Parse(strings.Split(strings.Split(accept, ",")[0], ";")[0])
			if err != nil {
				next.ServeHTTP(res, req)
				return
			}
			next.ServeHTTP(res, req.WithContext(i18next.SetLang(req.Context(), lang)))
		})
	}
}
