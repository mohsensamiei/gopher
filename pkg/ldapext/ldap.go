package ldapext

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/mohsensamiei/gopher/v3/pkg/authenticate"
	"github.com/mohsensamiei/gopher/v3/pkg/authorize"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/mapper"
	"google.golang.org/grpc/codes"
	"strings"
)

type LdapConfigs struct {
	Base Configs
}

func NewLdap(configs LdapConfigs) (*Ldap, error) {
	client, err := Open(configs.Base)
	if err != nil {
		return nil, err
	}
	return &Ldap{
		Client:      client,
		LdapConfigs: configs,
	}, nil
}

type Ldap struct {
	LdapConfigs
	Client *Client
}

func (a Ldap) Authorize(auth authenticate.Authenticate, scopes ...string) (*authorize.Claims, error) {
	switch cred := auth.(type) {
	case *authenticate.Basic:
		if a.Base.LdapBindUsername == cred.Username && a.Base.LdapBindPassword == cred.Password {
			return &authorize.Claims{
				ID:       fmt.Sprintf("cn=%v,%v", a.Base.LdapBindUsername, a.Base.LdapBaseDN),
				Username: a.Base.LdapBindUsername,
			}, nil
		}

		filters := []string{
			fmt.Sprintf("(cn=%v)", cred.Username),
		}
		if len(scopes) > 0 {
			filters = append(filters, fmt.Sprintf("(%v)", strings.Join(scopes, ",")))
		}
		users, _, err := a.Client.Users.Search(filters...)
		if err != nil {
			return nil, err
		}
		for _, user := range users {
			if err = a.check(user, cred.Password); err != nil {
				continue
			}
			claim := new(authorize.Claims)
			mapper.Struct(user, claim)
			return claim, nil
		}
	}
	return nil, errors.New(codes.Unauthenticated)
}

func (a Ldap) check(user *User, password string) error {
	conn, err := ldap.DialURL(a.Base.LdapURL)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()
	if err = conn.Bind(user.ID, password); err != nil {
		return err
	}
	return nil
}
