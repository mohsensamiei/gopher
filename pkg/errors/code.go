package errors

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
)

func Code(err error) codes.Code {
	if e := Cast(err); e != nil {
		return e.Code()
	}
	return codes.OK
}

func IsHandledCode(code codes.Code) bool {
	if status := runtime.HTTPStatusFromCode(code); status >= 500 && status <= 599 {
		return false
	}
	return true
}
