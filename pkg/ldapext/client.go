package ldapext

import "github.com/go-ldap/ldap/v3"

type Client struct {
	BaseDN string
	BindDN string
	Conn   *ldap.Conn
	Users  *Users
	Units  *Units
	Groups *Groups
}

func (c *Client) Close() {
	c.Conn.Close()
}
