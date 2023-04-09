package errors

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
	return getStatusLocalizedMessage(e.st).Locale
}

func (e Error) Localize() string {
	return getStatusLocalizedMessage(e.st).Message
}

func (e Error) SetLocalize(localize string) *Error {
	st := newStatusWithLocalize(e.Code(), e.Message(), e.Slug(), localize)
	st, _ = st.WithDetails(getStatusDetails(e.st)...)
	return &Error{st}
}

func (e Error) Details() []string {
	var res []string
	for _, ei := range getStatusErrorInfo(e.st) {
		res = append(res, ei.Reason)
	}
	return res
}

func (e Error) Validations() []*Validation {
	var res []*Validation
	for _, br := range getStatusBadRequest(e.st) {
		for _, v := range br.FieldViolations {
			res = append(res, &Validation{
				Tag:   v.Description,
				Field: v.Field,
			})
		}
	}
	return res
}

func (e Error) Error() string {
	parts := []string{
		strings.ToUpper(e.st.Code().String()),
		e.Message(),
		e.Slug(),
	}
	parts = append(parts, e.Details()...)
	for _, validation := range e.Validations() {
		parts = append(parts, fmt.Sprintf("tag '%v' on '%v'",
			validation.Tag,
			validation.Field,
		))
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
	var ei []proto.Message
	for _, detail := range details {
		ei = append(ei, &errdetails.ErrorInfo{
			Reason: detail,
		})
	}
	st, _ := e.st.WithDetails(ei...)
	return &Error{st}
}

func (e Error) WithValidation(tag string, field string) *Error {
	return e.WithValidations([]*Validation{
		{
			Tag:   tag,
			Field: field,
		},
	})
}

func (e Error) WithValidations(validations []*Validation) *Error {
	br := new(errdetails.BadRequest)
	for _, v := range validations {
		br.FieldViolations = append(br.FieldViolations, &errdetails.BadRequest_FieldViolation{
			Field:       v.Field,
			Description: v.Tag,
		})
	}
	st, _ := e.st.WithDetails(br)
	return &Error{st}
}
