package errors

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Wrap(err error, code codes.Code) *Error {
	return WrapWithSlug(err, code, strcase.ToSnake(code.String()))
}

func WrapWithSlug(err error, code codes.Code, slug string) *Error {
	var wst *status.Status
	switch e := err.(type) {
	case grpcStatus:
		wst, _ = newStatus(code, e.GRPCStatus().Message(), slug).WithDetails(getStatusDetails(e.GRPCStatus())...)
	default:
		wst = newStatus(code, err.Error(), slug)
	}
	return &Error{wst}
}
