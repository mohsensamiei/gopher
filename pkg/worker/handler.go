package worker

import (
	"github.com/pinosell/gopher/pkg/httpext"
	"net/http"
)

type Handler struct {
}

func (Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	httpext.SendCode(res, req, http.StatusOK)
}
