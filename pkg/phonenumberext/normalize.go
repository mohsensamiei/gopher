package phonenumberext

import (
	"github.com/dongri/phonenumber"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/nyaruka/phonenumbers"
	"google.golang.org/grpc/codes"
	"strings"
)

func Normalize(mobile string) (string, error) {
	mobile = strings.TrimLeft(mobile, "00")
	normal := phonenumbers.NormalizeDigitsOnly(mobile)
	alpha2 := phonenumber.GetISO3166ByNumber(normal, false).Alpha2
	if _, err := phonenumbers.Parse(normal, alpha2); err != nil {
		return "", errors.Wrap(err, codes.InvalidArgument)
	}
	return normal, nil
}
