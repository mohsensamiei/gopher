package authorize

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/authenticate"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
)

func Authorized(ctx context.Context, token authenticate.Authenticate, scopes ...string) (*Claims, error) {
	claim, err := di.Provide[Authorize](ctx).Authorize(token, scopes...)
	if err != nil {
		return nil, err
	}
	return claim, nil
}
