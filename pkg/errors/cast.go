package errors

import (
	"fmt"
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
		if local := statusLocalize(st); local == nil {
			return Wrap(fmt.Errorf(st.Message()), st.Code()).WithDetails(statusInfo(st)...)
		} else {
			return &Error{st: st}
		}
	}
	return Wrap(err, codes.Unknown)
}
