package authorize

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/pinosell/gopher/pkg/authenticate"
	"github.com/pinosell/gopher/pkg/errors"
	"github.com/pinosell/gopher/pkg/ldapext"
	"github.com/pinosell/gopher/pkg/mapper"
	"google.golang.org/grpc/codes"
)

type LdapConfigs struct {
	Base       ldapext.Configs
	LdapAdmins string `env:"LDAP_ADMINS,required"`
}

func NewLdap(configs LdapConfigs) (*Ldap, error) {
	client, err := ldapext.Open(configs.Base)
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
	Client *ldapext.Client
}

func (a Ldap) Authorize(auth authenticate.Authenticate, needAdmin bool) (*Claims, error) {
	switch cred := auth.(type) {
	case *authenticate.Basic:
		if a.Base.LdapBindUsername == cred.Username && a.Base.LdapBindPassword == cred.Password {
			return &Claims{
				UserID:   fmt.Sprintf("cn=%v,%v", a.Base.LdapBindUsername, a.Base.LdapBaseDN),
				Username: a.Base.LdapBindUsername,
			}, nil
		}

		filters := []string{
			fmt.Sprintf("(cn=%v)", cred.Username),
		}
		if needAdmin {
			filters = append(filters, fmt.Sprintf("(%v)", a.LdapAdmins))
		}
		users, _, err := a.Client.Users.Search(filters...)
		if err != nil {
			return nil, err
		}
		for _, user := range users {
			if err = a.check(user, cred.Password); err != nil {
				continue
			}
			claim := new(Claims)
			mapper.Struct(user, claim)
			return claim, nil
		}
	}
	return nil, errors.New(codes.Unauthenticated)
}

func (a Ldap) check(user *ldapext.User, password string) error {
	conn, err := ldap.DialURL(a.Base.LdapURL)
	if err != nil {
		return err
	}
	defer conn.Close()
	if err = conn.Bind(user.ID, password); err != nil {
		return err
	}
	return nil
}
