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

func getStatusLocalizedMessage(st *status.Status) *errdetails.LocalizedMessage {
	for _, detail := range st.Details() {
		if ei, ok := detail.(*errdetails.LocalizedMessage); ok {
			return ei
		}
	}
	return nil
}

func getStatusDetails(st *status.Status) []proto.Message {
	var res []proto.Message
	for _, detail := range st.Details() {
		dt, ok := detail.(proto.Message)
		if !ok {
			continue
		}
		res = append(res, dt)
	}
	return res
}

func getStatusErrorInfo(st *status.Status) []*errdetails.ErrorInfo {
	var details []*errdetails.ErrorInfo
	for _, detail := range st.Details() {
		ei, ok := detail.(*errdetails.ErrorInfo)
		if !ok {
			continue
		}
		details = append(details, ei)
	}
	return details
}

func getStatusBadRequest(st *status.Status) []*errdetails.BadRequest {
	var details []*errdetails.BadRequest
	for _, detail := range st.Details() {
		br, ok := detail.(*errdetails.BadRequest)
		if !ok {
			continue
		}
		details = append(details, br)
	}
	return details
}
