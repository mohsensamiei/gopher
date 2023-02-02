package metadataext

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func SetValue(ctx context.Context, key, value string) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, key, value)
	ctx = context.WithValue(ctx, key, value)
	return ctx
}

func GetValue(ctx context.Context, key string) (string, bool) {
	value, ok := ctx.Value(key).(string)
	if ok {
		return value, true
	}
	var md metadata.MD
	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	if len(md[key]) > 0 {
		return md[key][0], true
	}
	return "", false
}
