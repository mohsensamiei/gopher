package health

import (
	"github.com/mohsensamiei/gopher/v2/pkg/httpext"
	"github.com/mohsensamiei/gopher/v2/pkg/mimeext"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func NewHandler() *handler {
	return new(handler)
}

type handler struct {
}

func (handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/ping":
		httpext.Send(res, req, http.StatusOK, mimeext.Text, []byte("pong"))
	case "/metrics":
		promhttp.Handler().ServeHTTP(res, req)
	default:
		httpext.Send(res, req, http.StatusNotFound, mimeext.Text, []byte("not found"))
	}
}
