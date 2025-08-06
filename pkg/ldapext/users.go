package ldapext

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/query"
	"github.com/mohsensamiei/gopher/v3/pkg/slices"
	"google.golang.org/grpc/codes"
	"strconv"
	"strings"
)

const (
	startUID      = 1000
	startGID      = 1000
	userStructure = "inetOrgPerson"
)

type Users struct {
	client *Client
}

func (e Users) setSystemIDs(user *User) error {
	searchReq := ldap.NewSearchRequest(e.client.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(objectClass=%v)", userStructure),
		[]string{userAttrs.UID, userAttrs.GID},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		return err
	}

	user.UID = startUID
	user.GID = startGID
	for _, entity := range result.Entries {
		if a := attrValue(userAttrs.UID, entity); a != nil {
			if v, _ := strconv.Atoi(a[0]); v >= user.UID {
				user.UID = v + 1
			}
		}
		if a := attrValue(userAttrs.GID, entity); a != nil {
			if v, _ := strconv.Atoi(a[0]); v >= user.GID {
				user.GID = v + 1
			}
		}
	}
	return nil
}

// Deprecated: This method will be removed soon, Use the SetPassword method instead
func (e Users) Password(user *User, oldPassword, newPassword string) error {
	if _, err := e.client.Conn.PasswordModify(ldap.NewPasswordModifyRequest(
		user.ID,
		oldPassword,
		newPassword,
	)); err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultUnwillingToPerform) {
			return errors.Wrap(err, codes.InvalidArgument).
				WithDetailF("invalid old password '%v'", oldPassword)
		}
		return err
	}
	return nil
}

func (e Users) SetPassword(user *User, password string) error {
	if _, err := e.client.Conn.PasswordModify(ldap.NewPasswordModifyRequest(
		user.ID, "", password,
	)); err != nil {
		return err
	}
	return nil
}

func (e Users) Member(user *User, groups []*Group) error {
	for _, group := range user.Groups {
		searchReq := ldap.NewSearchRequest(group.ID,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			fmt.Sprintf("(objectClass=%v)", groupStructure),
			nil,
			nil)
		result, err := e.client.Conn.Search(searchReq)
		if err != nil {
			return err
		}
		for _, entry := range result.Entries {
			members := attrValue("uniqueMember", entry)
			members = slices.Remove(members, user.ID)

			req := ldap.NewModifyRequest(group.ID, nil)
			req.Replace("uniqueMember", members)

			if err = e.client.Conn.Modify(req); err != nil {
				return err
			}
		}
	}
	for _, group := range groups {
		searchReq := ldap.NewSearchRequest(group.ID,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			fmt.Sprintf("(objectClass=%v)", groupStructure),
			nil,
			nil)
		result, err := e.client.Conn.Search(searchReq)
		if err != nil {
			return err
		}
		for _, entry := range result.Entries {
			members := attrValue("uniqueMember", entry)
			members = append(members, user.ID)

			req := ldap.NewModifyRequest(group.ID, nil)
			req.Replace("uniqueMember", members)

			if err = e.client.Conn.Modify(req); err != nil {
				return err
			}
		}
	}
	return nil
}

func (e Users) Delete(user *User) error {
	if err := e.client.Conn.Del(ldap.NewDelRequest(
		user.ID,
		nil,
	)); err != nil {
		return err
	}
	return nil
}

func (e Users) Create(user *User, password string) error {
	unit, err := e.client.Units.Return(user.UnitID, query.Empty)
	if err != nil {
		return err
	}

	switch _, err = e.ReturnByUsername(user.Username, query.Empty); errors.Code(err) {
	case codes.NotFound:
	case codes.OK:
		return errors.New(codes.AlreadyExists).
			WithDetailF("there is another user with username '%v' exists", user.Username)
	default:
		return err
	}

	user.ID = fmt.Sprintf("cn=%v,%v", user.Username, unit.ID)
	if err = e.setSystemIDs(user); err != nil {
		return err
	}

	addReq := ldap.NewAddRequest(user.ID, []ldap.Control{})
	addReq.Attribute("objectClass", []string{userStructure, "posixAccount", "shadowAccount"})
	addReq.Attributes = append(addReq.Attributes, user.attrs()...)
	if err = e.client.Conn.Add(addReq); err != nil {
		return err
	}

	if err = e.SetPassword(user, password); err != nil {
		return err
	}
	return nil
}

func (e Users) List(unit *Unit, qe query.Encode) ([]*User, int64, error) {
	searchReq := ldap.NewSearchRequest(unit.ID,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(objectClass=%v)", userStructure),
		[]string{"*", "+"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		return nil, 0, err
	}

	var list []*User
	for _, entry := range result.Entries {
		user := &User{
			Unit: unit,
		}
		user.Groups, err = e.client.Groups.Returns(attrValue("memberOf", entry)...)
		if err != nil {
			return nil, 0, err
		}
		user.parse(entry)
		list = append(list, user)
	}
	return list, int64(len(list)), nil
}

func (e Users) Search(filters ...string) ([]*User, int64, error) {
	filters = append([]string{fmt.Sprintf("(objectClass=%v)", userStructure)}, filters...)
	searchReq := ldap.NewSearchRequest(e.client.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&%v)", strings.Join(filters, "")),
		[]string{"*", "+"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		return nil, 0, err
	}

	var list []*User
	for _, entry := range result.Entries {
		user := new(User)
		user.parse(entry)
		user.Unit, err = e.client.Units.Return(user.UnitID, query.Empty)
		if err != nil {
			return nil, 0, err
		}
		user.Groups, err = e.client.Groups.Returns(attrValue("memberOf", entry)...)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, user)
	}
	return list, int64(len(list)), nil
}

func (e Users) ReturnByUsername(username string, qe query.Encode) (*User, error) {
	searchReq := ldap.NewSearchRequest(e.client.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=%v)(cn=%v))", userStructure, username),
		[]string{"*", "+"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultNoSuchObject) {
			return nil, errors.Wrap(err, codes.NotFound).
				WithDetailF("there is user with username '%v' does not exists", username)
		}
		return nil, err
	}

	if len(result.Entries) <= 0 {
		return nil, errors.New(codes.NotFound).
			WithDetailF("there is user with username '%v' does not exists", username)
	}
	entry := result.Entries[0]

	user := new(User)
	user.parse(entry)
	user.Unit, err = e.client.Units.Return(user.UnitID, query.Empty)
	if err != nil {
		return nil, err
	}
	user.Groups, err = e.client.Groups.Returns(attrValue("memberOf", entry)...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e Users) Return(id string, qe query.Encode) (*User, error) {
	searchReq := ldap.NewSearchRequest(id,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(objectClass=%v)", userStructure),
		[]string{"*", "+"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultNoSuchObject) {
			return nil, errors.Wrap(err, codes.NotFound).
				WithDetailF("there is user with id '%v' does not exists", id)
		}
		return nil, err
	}

	if len(result.Entries) <= 0 {
		return nil, errors.Wrap(err, codes.NotFound).
			WithDetailF("there is user with id '%v' does not exists", id)
	}
	entry := result.Entries[0]

	user := new(User)
	user.parse(entry)
	user.Unit, err = e.client.Units.Return(user.UnitID, query.Empty)
	if err != nil {
		return nil, err
	}
	user.Groups, err = e.client.Groups.Returns(attrValue("memberOf", entry)...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e Users) Update(user *User) error {
	req := ldap.NewModifyRequest(user.ID, nil)
	for _, attribute := range user.attrs() {
		req.Replace(attribute.Type, attribute.Vals)
	}
	if err := e.client.Conn.Modify(req); err != nil {
		return err
	}
	return nil
}
