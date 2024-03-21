package authorize

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/authenticate"
	"github.com/mohsensamiei/gopher/pkg/di"
)

func Authorized(ctx context.Context, token authenticate.Authenticate, scopes ...string) (*Claims, error) {
	claim, err := di.Provide[Authorize](ctx, Name).Authorize(token, scopes...)
	if err != nil {
		return nil, err
	}
	return claim, nil
}
