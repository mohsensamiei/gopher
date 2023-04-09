package phonenumberext

import (
	"fmt"
	"github.com/dongri/phonenumber"
	"github.com/nyaruka/phonenumbers"
	"github.com/pinosell/gopher/pkg/errors"
	"google.golang.org/grpc/codes"
	"strings"
)

func Normalize(mobile string) (string, error) {
	mobile = strings.TrimLeft(mobile, "00")
	normal := phonenumbers.NormalizeDigitsOnly(mobile)
	alpha2 := phonenumber.GetISO3166ByNumber(normal, false).Alpha2
	num, err := phonenumbers.Parse(normal, alpha2)
	if err != nil {
		return "", errors.Wrap(err, codes.InvalidArgument)
	}
	return fmt.Sprintf("%v", *num.NationalNumber), nil
}
