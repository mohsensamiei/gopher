package jwtext

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohsensamiei/gopher/v2/pkg/authenticate"
	"github.com/mohsensamiei/gopher/v2/pkg/authorize"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/mohsensamiei/gopher/v2/pkg/slices"
	"google.golang.org/grpc/codes"
	"time"
)

type JWT struct {
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
	duration  time.Duration
}

type Claims struct {
	*authorize.Claims
	*jwt.RegisteredClaims
}

func (c Claims) Valid() error {
	now := time.Now().Unix()
	if c.ExpiresAt.Unix() < now || c.IssuedAt.Unix() > now {
		return errors.New(codes.DeadlineExceeded)
	}
	return nil
}

func (a JWT) Authorize(auth authenticate.Authenticate, scopes ...string) (*authorize.Claims, error) {
	switch cred := auth.(type) {
	case *authenticate.Bearer:
		at, err := jwt.ParseWithClaims(cred.Token, &Claims{}, func(_ *jwt.Token) (any, error) {
			return a.verifyKey, nil
		})
		if err != nil {
			return nil, errors.Wrap(err, codes.Unauthenticated)
		}
		claims := at.Claims.(*Claims)
		if slices.Common(claims.Scopes, scopes, 1) {
			return nil, errors.New(codes.PermissionDenied)
		}
		return claims.Claims, nil
	}
	return nil, errors.New(codes.Unauthenticated)
}

func (a JWT) Sign(c *Claims) (string, error) {
	now := time.Now()
	c.IssuedAt = &jwt.NumericDate{Time: now}
	c.ExpiresAt = &jwt.NumericDate{Time: now.Add(a.duration)}

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	token, err := at.SignedString(a.signKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
