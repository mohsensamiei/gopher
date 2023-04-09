package phonenumberext

import (
	"github.com/dongri/phonenumber"
	"github.com/nyaruka/phonenumbers"
)

func CountryAlpha2(mobile string) string {
	normal := phonenumbers.NormalizeDigitsOnly(mobile)
	return phonenumber.GetISO3166ByNumber(normal, false).Alpha2
}
