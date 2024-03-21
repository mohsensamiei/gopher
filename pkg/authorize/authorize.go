package authorize

import (
	"github.com/mohsensamiei/gopher/pkg/authenticate"
)

type Authorize interface {
	Authorize(auth authenticate.Authenticate, scopes ...string) (*Claims, error)
}
