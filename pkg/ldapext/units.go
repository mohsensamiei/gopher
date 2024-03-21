package ldapext

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/mohsensamiei/gopher/pkg/errors"
	"github.com/mohsensamiei/gopher/pkg/query"
	"google.golang.org/grpc/codes"
)

const (
	unitStructure = "organizationalUnit"
)

type Units struct {
	client *Client
}

func (e Units) Delete(unit *Unit) error {
	searchReq := ldap.NewSearchRequest(unit.ID,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(!(objectClass=%v))", unitStructure),
		[]string{"dn"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		return err
	}
	for _, entry := range result.Entries {
		if err = e.client.Conn.Del(ldap.NewDelRequest(
			entry.DN,
			nil,
		)); err != nil {
			return err
		}
	}

	if err = e.client.Conn.Del(ldap.NewDelRequest(
		unit.ID,
		nil,
	)); err != nil {
		return err
	}
	return nil
}

func (e Units) Create(unit *Unit) error {
	unit.ID = fmt.Sprintf("ou=%v,%v", unit.Name, e.client.BaseDN)
	if _, err := e.Return(unit.ID, query.Empty); err == nil {
		return errors.New(codes.AlreadyExists).
			WithDetailF("there is another unit with id '%v' exists", unit.ID)
	}

	addReq := ldap.NewAddRequest(unit.ID, []ldap.Control{})
	addReq.Attribute("objectClass", []string{unitStructure})
	addReq.Attributes = append(addReq.Attributes, unit.attrs()...)

	if err := e.client.Conn.Add(addReq); err != nil {
		return err
	}
	return nil
}

func (e Units) List(qe query.Encode) ([]*Unit, int64, error) {
	searchReq := ldap.NewSearchRequest(e.client.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(objectClass=%v)", unitStructure),
		[]string{"*", "+"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		return nil, 0, err
	}

	var list []*Unit
	for _, entry := range result.Entries {
		unit := new(Unit)
		unit.parse(entry)
		list = append(list, unit)
	}
	return list, int64(len(list)), nil
}

func (e Units) Return(id string, qe query.Encode) (*Unit, error) {
	searchReq := ldap.NewSearchRequest(id,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(objectClass=%v)", unitStructure),
		[]string{"*", "+"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultNoSuchObject) {
			return nil, errors.Wrap(err, codes.NotFound).
				WithDetailF("there is unit with id '%v' does not exists", id)
		}
		return nil, err
	}

	if len(result.Entries) <= 0 {
		return nil, errors.New(codes.NotFound).
			WithDetailF("there is unit with id '%v' does not exists", id)
	}
	entry := result.Entries[0]

	unit := new(Unit)
	unit.parse(entry)
	return unit, nil
}
