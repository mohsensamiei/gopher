package httpext

import (
	"github.com/mohsensamiei/gopher/v3/pkg/authenticate"
	"net/http"
)

func AuthRequest(req *http.Request) (authenticate.Authenticate, error) {
	return authenticate.Decode(req.Header.Get(AuthorizationHeader))
}
