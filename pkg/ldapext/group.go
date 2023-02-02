package ldapext

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"strings"
	"time"
)

var (
	groupAttrs = struct {
		Name      string
		CreatedAt string
		UpdatedAt string
	}{
		Name:      "cn",
		CreatedAt: "createTimestamp",
		UpdatedAt: "modifyTimestamp",
	}
)

type Group struct {
	ID        string
	UnitID    string
	Unit      *Unit
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Group) attrs() (attrs []ldap.Attribute) {
	attrs = append(attrs, ldap.Attribute{Type: groupAttrs.Name, Vals: []string{e.Name}})
	return
}

func (e *Group) parse(entry *ldap.Entry) {
	e.ID = entry.DN
	for _, attr := range entry.Attributes {
		switch attr.Name {
		case groupAttrs.Name:
			e.Name = attr.Values[0]
		case groupAttrs.CreatedAt:
			e.CreatedAt, _ = time.Parse(timestampFormat, attr.Values[0])
		case groupAttrs.UpdatedAt:
			e.UpdatedAt, _ = time.Parse(timestampFormat, attr.Values[0])
		}
	}
	e.UnitID = strings.ReplaceAll(entry.DN, fmt.Sprintf("cn=%v,", e.Name), "")
}
