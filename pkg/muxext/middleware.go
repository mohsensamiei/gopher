package muxext

import (
	"github.com/gorilla/mux"
	"github.com/mohsensamiei/gopher/v3/pkg/authorize"
	"github.com/mohsensamiei/gopher/v3/pkg/cache"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"github.com/mohsensamiei/gopher/v3/pkg/httpext"
	"github.com/mohsensamiei/gopher/v3/pkg/i18next"
	"golang.org/x/text/language"
	"net/http"
)

func AuthMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			_ = httpRequestAuth(req)
			next.ServeHTTP(res, req)
		})
	}
}

func httpRequestAuth(req *http.Request) error {
	token, err := httpext.AuthRequest(req)
	if err != nil {
		return err
	}
	if _, err = authorize.Authorized(req.Context(), token); err != nil {
		return err
	}
	*req = *req.WithContext(authorize.ToContext(req.Context(), token))
	return nil
}

func DIMiddleware[T any](provide func() T) mux.MiddlewareFunc {
	value := provide()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(res, req.WithContext(di.Register(req.Context(), value)))
		})
	}
}

func LangMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			accept := httpext.Header(req, httpext.AcceptLanguageHeader)
			if accept == "" {
				next.ServeHTTP(res, req)
				return
			}
			var base language.Base
			if tags, _, err := language.ParseAcceptLanguage(accept); err != nil {
				next.ServeHTTP(res, req)
				return
			} else {
				base, _ = tags[0].Base()
			}
			lang, err := language.Parse(base.String())
			if err != nil {
				next.ServeHTTP(res, req)
				return
			}
			next.ServeHTTP(res, req.WithContext(i18next.SetLang(req.Context(), lang)))
		})
	}
}

func CacheMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			di.Provide[*cache.Client](req.Context()).Middleware(next).ServeHTTP(res, req)
		})
	}
}
