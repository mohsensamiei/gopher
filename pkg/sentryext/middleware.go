package sentryext

import (
	sentryhttp "github.com/getsentry/sentry-go/http"
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	return sentryhttp.New(sentryhttp.Options{
		Repanic: false,
	}).Handle(handler)
}
