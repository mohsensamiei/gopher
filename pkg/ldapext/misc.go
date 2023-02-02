package ldapext

import (
	"github.com/go-ldap/ldap/v3"
)

const (
	timestampFormat = "20060102150405Z"
)

func attrValue(name string, entry *ldap.Entry) []string {
	for _, attribute := range entry.Attributes {
		if attribute.Name != name {
			continue
		}
		return attribute.Values
	}
	return nil
}
