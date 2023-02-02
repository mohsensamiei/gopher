package ldapext

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"strconv"
	"strings"
	"time"
)

var (
	userAttrs = struct {
		Email         string
		Mobile        string
		Surname       string
		Username      string
		Name          string
		UID           string
		GID           string
		HomeDirectory string
		CreatedAt     string
		UpdatedAt     string
	}{
		Email:         "mail",
		Mobile:        "mobile",
		Surname:       "sn",
		Username:      "uid",
		Name:          "givenName",
		UID:           "uidNumber",
		GID:           "gidNumber",
		HomeDirectory: "homeDirectory",
		CreatedAt:     "createTimestamp",
		UpdatedAt:     "modifyTimestamp",
	}
)

type User struct {
	ID        string
	UnitID    string
	Unit      *Unit
	Username  string
	Name      string
	Surname   string
	Email     string
	Mobile    string
	UID       int
	GID       int
	Groups    []*Group
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *User) attrs() (attrs []ldap.Attribute) {
	attrs = append(attrs, ldap.Attribute{Type: userAttrs.Email, Vals: []string{e.Email}})
	attrs = append(attrs, ldap.Attribute{Type: userAttrs.Mobile, Vals: []string{e.Mobile}})
	attrs = append(attrs, ldap.Attribute{Type: userAttrs.Surname, Vals: []string{e.Surname}})
	attrs = append(attrs, ldap.Attribute{Type: userAttrs.Username, Vals: []string{e.Username}})
	attrs = append(attrs, ldap.Attribute{Type: userAttrs.Name, Vals: []string{e.Name}})
	attrs = append(attrs, ldap.Attribute{Type: userAttrs.UID, Vals: []string{fmt.Sprint(e.UID)}})
	attrs = append(attrs, ldap.Attribute{Type: userAttrs.GID, Vals: []string{fmt.Sprint(e.GID)}})
	attrs = append(attrs, ldap.Attribute{Type: userAttrs.HomeDirectory, Vals: []string{fmt.Sprintf("/home/%v", e.Username)}})
	return
}

func (e *User) parse(entry *ldap.Entry) {
	e.ID = entry.DN
	for _, attr := range entry.Attributes {
		switch attr.Name {
		case userAttrs.UID:
			e.UID, _ = strconv.Atoi(attr.Values[0])
		case userAttrs.GID:
			e.GID, _ = strconv.Atoi(attr.Values[0])
		case userAttrs.Name:
			e.Name = attr.Values[0]
		case userAttrs.Surname:
			e.Surname = attr.Values[0]
		case userAttrs.Username:
			e.Username = attr.Values[0]
		case userAttrs.Email:
			e.Email = attr.Values[0]
		case userAttrs.Mobile:
			e.Mobile = attr.Values[0]
		case userAttrs.CreatedAt:
			e.CreatedAt, _ = time.Parse(timestampFormat, attr.Values[0])
		case userAttrs.UpdatedAt:
			e.UpdatedAt, _ = time.Parse(timestampFormat, attr.Values[0])
		}
		e.UnitID = strings.ReplaceAll(entry.DN, fmt.Sprintf("cn=%v,", e.Username), "")
	}
}
