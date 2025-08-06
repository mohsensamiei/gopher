package recovery

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (res any, err error) {
		defer func() {
			if rec := recover(); rec != nil {
				err = handlePanic(rec)
			}
			if !errors.IsHandledCode(errors.Code(err)) {
				log.WithError(err).
					WithField("method", info.FullMethod).
					Error("unhandled error")
			}
		}()
		res, err = handler(ctx, req)
		return
	}
}
