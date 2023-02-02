package errors

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/grpc/codes"
)

func Wrap(err error, code codes.Code) *Error {
	return WrapWithSlug(err, code, strcase.ToSnake(code.String()))
}

func WrapWithSlug(err error, code codes.Code, slug string) *Error {
	return &Error{newStatus(code, err.Error(), slug)}
}
