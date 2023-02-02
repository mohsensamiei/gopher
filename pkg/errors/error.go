package errors

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Error struct {
	st *status.Status
}

func (e Error) Message() string {
	return e.st.Message()
}

func (e Error) Slug() string {
	return statusLocalize(e.st).Locale
}

func (e Error) Localize() string {
	return statusLocalize(e.st).Message
}

func (e Error) SetLocalize(localize string) *Error {
	st := newStatusWithLocalize(e.Code(), e.Message(), e.Slug(), localize)
	st = statusWithDetails(st, e.Details())
	return &Error{st}
}

func (e Error) Details() []string {
	return statusInfo(e.st)
}

func (e Error) Error() string {
	parts := []string{
		strings.ToUpper(e.st.Code().String()),
		e.Message(),
		e.Slug(),
	}
	if len(e.Details()) > 0 {
		parts = append(parts, strings.Join(e.Details(), ", "))
	}
	return strings.Join(parts, " | ")
}

func (e Error) GRPCStatus() *status.Status {
	return e.st
}

func (e Error) Code() codes.Code {
	return e.st.Code()
}

func (e Error) StatusCode() int {
	return runtime.HTTPStatusFromCode(e.st.Code())
}

func (e Error) IsHandled() bool {
	return IsHandledCode(e.st.Code())
}

func (e Error) WithDetailF(format string, args ...any) *Error {
	return e.WithDetails(fmt.Sprintf(format, args...))
}

func (e Error) WithDetails(details ...string) *Error {
	return &Error{statusWithDetails(e.st, details)}
}
