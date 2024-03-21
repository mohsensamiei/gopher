package worker

import (
	"github.com/mohsensamiei/gopher/pkg/httpext"
	"net/http"
)

type Handler struct {
}

func (Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	httpext.SendCode(res, req, http.StatusOK)
}
