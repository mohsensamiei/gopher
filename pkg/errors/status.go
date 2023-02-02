package errors

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newStatus(code codes.Code, message, slug string) *status.Status {
	return newStatusWithLocalize(code, message, slug, "")
}

func newStatusWithLocalize(code codes.Code, message, slug string, localize string) *status.Status {
	st, _ := status.New(code, message).WithDetails(&errdetails.LocalizedMessage{
		Locale:  slug,
		Message: localize,
	})
	return st
}

func statusLocalize(st *status.Status) *errdetails.LocalizedMessage {
	for _, detail := range st.Details() {
		if ei, ok := detail.(*errdetails.LocalizedMessage); ok {
			return ei
		}
	}
	return nil
}

func statusInfo(st *status.Status) []string {
	var details []string
	for _, detail := range st.Details() {
		ei, ok := detail.(*errdetails.ErrorInfo)
		if !ok {
			continue
		}
		details = append(details, ei.Reason)
	}
	return details
}

func statusWithDetails(st *status.Status, details []string) *status.Status {
	var messages []proto.Message
	for _, detail := range details {
		messages = append(messages, &errdetails.ErrorInfo{
			Reason: detail,
		})
	}
	s, _ := st.WithDetails(messages...)
	return s
}
