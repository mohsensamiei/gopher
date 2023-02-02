package ldapext

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

type Configs struct {
	LdapURL          string `env:"LDAP_URL,required"`
	LdapBaseDN       string `env:"LDAP_BASE_DN,required"`
	LdapBindUsername string `env:"LDAP_BIND_USERNAME,required"`
	LdapBindPassword string `env:"LDAP_BIND_PASSWORD,required"`
}

func Open(configs Configs) (*Client, error) {
	conn, err := ldap.DialURL(configs.LdapURL)
	if err != nil {
		return nil, err
	}
	bindDN := fmt.Sprintf("cn=%v,%v",
		configs.LdapBindUsername,
		configs.LdapBaseDN,
	)
	if err = conn.Bind(bindDN, configs.LdapBindPassword); err != nil {
		return nil, err
	}

	client := &Client{
		Conn:   conn,
		BindDN: bindDN,
		BaseDN: configs.LdapBaseDN,
	}
	client.Units = &Units{client: client}
	client.Users = &Users{client: client}
	client.Groups = &Groups{client: client}
	return client, nil
}
