package ldapext

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/mohsensamiei/gopher/v2/pkg/query"
	"google.golang.org/grpc/codes"
)

const (
	groupStructure = "groupOfUniqueNames"
)

type Groups struct {
	client *Client
}

func (e Groups) Delete(group *Group) error {
	if err := e.client.Conn.Del(ldap.NewDelRequest(
		group.ID,
		nil,
	)); err != nil {
		return err
	}
	return nil
}

func (e Groups) Create(group *Group) error {
	unit, err := e.client.Units.Return(group.UnitID, query.Empty)
	if err != nil {
		return err
	}
	group.ID = fmt.Sprintf("cn=%v,%v", group.Name, unit.ID)
	if _, err = e.Return(group.ID, query.Empty); err == nil {
		return errors.New(codes.AlreadyExists).
			WithDetailF("there is another group with id '%v' exists", group.ID)
	}

	addReq := ldap.NewAddRequest(group.ID, []ldap.Control{})
	addReq.Attribute("objectClass", []string{groupStructure})
	addReq.Attribute("uniqueMember", []string{e.client.BindDN})
	addReq.Attributes = append(addReq.Attributes, group.attrs()...)

	if err = e.client.Conn.Add(addReq); err != nil {
		return err
	}
	return nil
}

func (e Groups) List(unit *Unit, _ query.Encode) ([]*Group, int64, error) {
	searchReq := ldap.NewSearchRequest(unit.ID,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(objectClass=%v)", groupStructure),
		[]string{"*", "+"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		return nil, 0, err
	}

	var list []*Group
	for _, entry := range result.Entries {
		group := &Group{
			Unit: unit,
		}
		group.parse(entry)
		list = append(list, group)
	}
	return list, int64(len(list)), nil
}

func (e Groups) Return(id string, _ query.Encode) (*Group, error) {
	searchReq := ldap.NewSearchRequest(id,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(objectClass=%v)", groupStructure),
		[]string{"*", "+"},
		nil)
	result, err := e.client.Conn.Search(searchReq)
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultNoSuchObject) {
			return nil, errors.Wrap(err, codes.NotFound).
				WithDetailF("there is group with id '%v' does not exists", id)
		}
		return nil, err
	}

	if len(result.Entries) <= 0 {
		return nil, errors.Wrap(err, codes.NotFound).
			WithDetailF("there is group with id '%v' does not exists", id)
	}
	entry := result.Entries[0]

	group := new(Group)
	group.parse(entry)
	group.Unit, err = e.client.Units.Return(group.UnitID, "")
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (e Groups) Returns(id ...string) ([]*Group, error) {
	var groups []*Group
	for _, i := range id {
		group, err := e.Return(i, "")
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
