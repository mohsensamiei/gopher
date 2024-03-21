package mysqlext

import (
	"fmt"

	"github.com/mohsensamiei/gopher/pkg/netext"
)

type Configs struct {
	MySQLHosts    string      `env:"MYSQL_HOSTS,required"`
	MySQLPort     netext.Port `env:"MYSQL_PORT" envDefault:"3306"`
	MySQLUsername string      `env:"MYSQL_USERNAME,required"`
	MySQLPassword string      `env:"MYSQL_PASSWORD,required"`
	MySQLDBName   string      `env:"MYSQL_DBNAME,required"`
}

func (Configs) DB() string {
	return "mysql"
}

func (c Configs) DSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		c.MySQLUsername,
		c.MySQLPassword,
		c.MySQLHosts,
		c.MySQLPort,
		c.MySQLDBName,
	)
}
