package errors

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcStatus interface {
	GRPCStatus() *status.Status
}

func Cast(err error) *Error {
	switch e := err.(type) {
	case nil:
		return nil
	case *Error:
		return e
	case grpcStatus:
		st := e.GRPCStatus()
		// it means status is not compatible with our package
		if local := getStatusLocalizedMessage(st); local == nil {
			wrapped, _ := newStatus(st.Code(), st.Message(), strcase.ToSnake(st.Code().String())).
				WithDetails(getStatusDetails(st)...)
			return &Error{st: wrapped}
		} else {
			return &Error{st: st}
		}
	}
	return Wrap(err, codes.Unknown)
}
