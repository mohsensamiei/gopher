package muxext

import (
	"github.com/gorilla/mux"
	"github.com/mohsensamiei/gopher/pkg/authorize"
	"github.com/mohsensamiei/gopher/pkg/di"
	"github.com/mohsensamiei/gopher/pkg/httpext"
	"github.com/mohsensamiei/gopher/pkg/i18next"
	"github.com/mohsensamiei/gopher/pkg/response"
	"golang.org/x/text/language"
	"net/http"
	"strings"
)

func AuthMiddleware(scopes ...string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			token, err := httpext.AuthHeader(req)
			if err != nil {
				httpext.SendError(res, req, err)
				return
			}
			var claims *authorize.Claims
			claims, err = authorize.Authorized(req.Context(), token, scopes...)
			if err != nil {
				httpext.SendError(res, req, err)
				return
			}
			*req = *req.WithContext(authorize.ToContext(req.Context(), claims))
			next.ServeHTTP(res, req)
		})
	}
}

func DIMiddleware[T any](tc di.ToContext[T], provide func() T) mux.MiddlewareFunc {
	value := provide()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(res, req.WithContext(tc(req.Context(), value)))
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
			lang, err := language.Parse(strings.Split(strings.Split(accept, ",")[0], ";")[0])
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
			response.FromContext(req.Context()).Middleware(next).ServeHTTP(res, req)
		})
	}
}
