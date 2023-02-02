package authorize

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func DIMiddleware(authorize Authorize) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx = context.WithValue(ctx, Name, authorize)
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}
