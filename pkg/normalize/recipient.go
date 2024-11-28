package normalize

import (
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"google.golang.org/grpc/codes"
)

func Recipient(val string) (string, error) {
	if normal, err := Phone(val); err == nil {
		return normal, nil
	}
	if normal, err := Email(val); err == nil {
		return normal, nil
	}
	return "", errors.NewWithSlug(codes.InvalidArgument, "invalid_recipient").
		WithDetails(val)
}
