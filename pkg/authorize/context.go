package authorize

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/authenticate"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"github.com/mohsensamiei/gopher/v3/pkg/metadataext"
)

func ToContext(ctx context.Context, auth authenticate.Authenticate) context.Context {
	return metadataext.SetValue(ctx, CurrentToken, authenticate.Encode(auth))
}

func TokenFromContext(ctx context.Context) (string, error) {
	token, ok := metadataext.GetValue(ctx, CurrentToken)
	if !ok {
		return "", authenticate.ErrUnauthenticated
	}
	auth, err := authenticate.Decode(token)
	if err != nil {
		return "", authenticate.ErrUnauthenticated
	}
	switch a := auth.(type) {
	case *authenticate.Bearer:
		return a.Token, nil
	default:
		return "", authenticate.ErrUnauthenticated
	}
}

func ClaimsFromContext(ctx context.Context) (*Claims, error) {
	token, ok := metadataext.GetValue(ctx, CurrentToken)
	if !ok {
		return nil, authenticate.ErrUnauthenticated
	}
	auth, err := authenticate.Decode(token)
	if err != nil {
		return nil, err
	}
	var claims *Claims
	claims, err = di.Provide[Authorize](ctx).Authorize(auth)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
