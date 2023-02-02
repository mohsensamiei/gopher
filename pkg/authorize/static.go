package authorize

import (
	"github.com/pinosell/gopher/pkg/authenticate"
	"github.com/pinosell/gopher/pkg/bcryptext"
	"github.com/pinosell/gopher/pkg/errors"
	"github.com/pinosell/gopher/pkg/slices"
	"google.golang.org/grpc/codes"
)

type StaticConfigs struct {
	StaticAdmins      []string          `env:"STATIC_ADMINS"`
	StaticCredentials map[string]string `env:"STATIC_CREDENTIALS"`
}

// Deprecated: This method will be removed soon, Use the NewLdap method instead
func NewStatic(configs StaticConfigs) *Static {
	return &Static{
		StaticConfigs: configs,
	}
}

type Static struct {
	StaticConfigs
}

func (a Static) Authorize(auth authenticate.Authenticate, needAdmin bool) (*Claims, error) {
	switch cred := auth.(type) {
	case *authenticate.Basic:
		password, ok := a.StaticCredentials[cred.Username]
		if !ok {
			break
		}
		if needAdmin && !slices.Contains(cred.Username, a.StaticAdmins...) {
			break
		}
		if err := bcryptext.CompareHashAndPassword(password, cred.Password); err == nil {
			return &Claims{
				Username: cred.Username,
			}, nil
		}
	}
	return nil, errors.New(codes.Unauthenticated)
}
