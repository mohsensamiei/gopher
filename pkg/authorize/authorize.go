package authorize

import (
	"github.com/pinosell/gopher/pkg/authenticate"
)

type Authorize interface {
	Authorize(auth authenticate.Authenticate, needAdmin bool) (*Claims, error)
}
