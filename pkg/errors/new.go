package errors

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/grpc/codes"
)

func New(code codes.Code) *Error {
	return NewWithSlug(code, strcase.ToSnake(code.String()))
}

func NewWithSlug(code codes.Code, slug string) *Error {
	return &Error{newStatus(code, strcase.ToCamel(code.String()), slug)}
}
