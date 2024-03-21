package httpext

import (
	"github.com/mohsensamiei/gopher/v2/pkg/authenticate"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"google.golang.org/grpc/codes"
	"net/http"
	"strings"
)

var (
	ErrInvalidToken = errors.New(codes.Unauthenticated).WithDetails("invalid Authenticate")
	tokenParsers    = map[authenticate.Type]func(req *http.Request) (authenticate.Authenticate, error){
		authenticate.BearerType: bearerParser,
		authenticate.BasicType:  basicParser,
	}
)

func bearerParser(req *http.Request) (authenticate.Authenticate, error) {
	if strings.HasPrefix(req.Header.Get(AuthorizationHeader), "Bearer") {
		return authenticate.NewBearer(strings.TrimSpace(strings.Replace(req.Header.Get(AuthorizationHeader), "Bearer ", "", 1))), nil
	}
	return nil, ErrInvalidToken
}

func basicParser(req *http.Request) (authenticate.Authenticate, error) {
	username, password, ok := req.BasicAuth()
	if ok {
		return authenticate.NewBasic(username, password), nil
	}
	return nil, ErrInvalidToken
}

func AuthHeader(req *http.Request) (authenticate.Authenticate, error) {
	for _, parser := range tokenParsers {
		if token, err := parser(req); err == nil {
			return token, nil
		}
	}
	return nil, errors.New(codes.Unauthenticated).WithDetails("no authenticate header")
}
