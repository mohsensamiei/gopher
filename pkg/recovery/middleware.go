package recovery

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AuthMiddleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					log.WithError(handlePanic(rec)).
						WithField("url", fmt.Sprint(req.Method, " ", req.RequestURI)).
						Error("unhandled error")
				}
			}()
			handler.ServeHTTP(res, req)
		})
	}
}
