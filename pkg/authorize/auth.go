package authorize

import (
	"github.com/gorilla/mux"
	"github.com/pinosell/gopher/pkg/di"
	"github.com/pinosell/gopher/pkg/httpext"
	"net/http"
)

func AuthHttpRequest(req *http.Request, needAdmin bool) error {
	token, err := httpext.AuthHeader(req)
	if err != nil {
		return err
	}

	var claim *Claims
	auth := di.Provide[Authorize](req.Context(), Name)
	if claim, err = auth.Authorize(token, needAdmin); err != nil {
		return err
	}

	*req = *req.WithContext(ToContext(req.Context(), claim))
	return nil
}

func AuthMiddleware(needAdmin bool) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if err := AuthHttpRequest(req, needAdmin); err != nil {
				httpext.SendError(res, req, err)
				return
			}
			next.ServeHTTP(res, req)
		})
	}
}
