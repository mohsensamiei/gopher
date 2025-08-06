package jwtext

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohsensamiei/gopher/v3/pkg/authenticate"
	"github.com/mohsensamiei/gopher/v3/pkg/authorize"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/slices"
	"google.golang.org/grpc/codes"
	"time"
)

type JWT struct {
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
	duration  time.Duration
}

type claims struct {
	*authorize.Claims
	*jwt.RegisteredClaims
}

func (c claims) Valid() error {
	now := time.Now().Unix()
	if c.ExpiresAt.Unix() < now || c.IssuedAt.Unix() > now {
		return errors.New(codes.DeadlineExceeded)
	}
	return nil
}

func (a JWT) Authorize(auth authenticate.Authenticate, scopes ...string) (*authorize.Claims, error) {
	switch cred := auth.(type) {
	case *authenticate.Bearer:
		at, err := jwt.ParseWithClaims(cred.Token, &claims{}, func(_ *jwt.Token) (any, error) {
			return a.verifyKey, nil
		})
		if err != nil {
			return nil, errors.Wrap(err, codes.Unauthenticated)
		}
		claim := at.Claims.(*claims)
		if scopes != nil && !slices.Common(claim.Scopes, scopes, 1) {
			return nil, errors.New(codes.PermissionDenied)
		}
		return claim.Claims, nil
	}
	return nil, errors.New(codes.Unauthenticated)
}

func (a JWT) Sign(c *authorize.Claims) (string, *jwt.RegisteredClaims, error) {
	claim := &claims{
		Claims:           c,
		RegisteredClaims: new(jwt.RegisteredClaims),
	}
	now := time.Now()
	claim.IssuedAt = &jwt.NumericDate{Time: now}
	claim.ExpiresAt = &jwt.NumericDate{Time: now.Add(a.duration)}

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	token, err := at.SignedString(a.signKey)
	if err != nil {
		return "", nil, err
	}
	return token, claim.RegisteredClaims, nil
}
