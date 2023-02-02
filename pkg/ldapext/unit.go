package ldapext

import (
	"github.com/go-ldap/ldap/v3"
	"time"
)

var (
	unitAttrs = struct {
		Name      string
		CreatedAt string
		UpdatedAt string
	}{
		Name:      "ou",
		CreatedAt: "createTimestamp",
		UpdatedAt: "modifyTimestamp",
	}
)

type Unit struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Unit) attrs() (attrs []ldap.Attribute) {
	attrs = append(attrs, ldap.Attribute{Type: unitAttrs.Name, Vals: []string{u.Name}})
	return
}

func (u *Unit) parse(entry *ldap.Entry) {
	u.ID = entry.DN
	for _, attr := range entry.Attributes {
		switch attr.Name {
		case unitAttrs.Name:
			u.Name = attr.Values[0]
		case unitAttrs.CreatedAt:
			u.CreatedAt, _ = time.Parse(timestampFormat, attr.Values[0])
		case unitAttrs.UpdatedAt:
			u.UpdatedAt, _ = time.Parse(timestampFormat, attr.Values[0])
		}
	}
}
