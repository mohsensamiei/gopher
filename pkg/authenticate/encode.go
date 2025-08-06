package authenticate

import (
	"encoding/base64"
	"fmt"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"google.golang.org/grpc/codes"
	"strings"
)

var (
	ErrUnauthenticated = errors.New(codes.Unauthenticated)
	parsers            = map[Type]func(string) (Authenticate, error){
		BearerType: bearerParser,
		BasicType:  basicParser,
	}
)

func bearerParser(auth string) (Authenticate, error) {
	prefix := fmt.Sprintf("%v ", BearerType)
	if !strings.HasPrefix(auth, prefix) {
		return nil, ErrUnauthenticated
	}
	_, token, ok := strings.Cut(auth, " ")
	if !ok {
		return nil, ErrUnauthenticated
	}
	return NewBearer(token), nil
}

func basicParser(auth string) (Authenticate, error) {
	prefix := fmt.Sprintf("%v ", BasicType)
	if !strings.HasPrefix(auth, prefix) {
		return nil, ErrUnauthenticated
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return nil, ErrUnauthenticated
	}
	username, password, ok := strings.Cut(string(c), ":")
	if !ok {
		return nil, ErrUnauthenticated
	}
	return NewBasic(username, password), nil
}

func Encode(auth Authenticate) string {
	switch t := auth.(type) {
	case *Bearer:
		return fmt.Sprintf("%v %v", BearerType, t.Token)
	case *Basic:
		return fmt.Sprintf("%v %v", BasicType,
			base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", t.Username, t.Password))))
	default:
		panic("unknown token type")
	}
}

func Decode(token string) (Authenticate, error) {
	for _, parser := range parsers {
		if auth, err := parser(token); err == nil {
			return auth, nil
		}
	}
	return nil, ErrUnauthenticated
}
