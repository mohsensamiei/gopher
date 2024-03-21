package postgresext

import (
	"fmt"
	"github.com/mohsensamiei/gopher/pkg/netext"
	"strings"
)

type Configs struct {
	PostgresHosts    []string    `env:"POSTGRES_HOSTS,required"`
	PostgresPort     netext.Port `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUsername string      `env:"POSTGRES_USERNAME,required"`
	PostgresPassword string      `env:"POSTGRES_PASSWORD,required"`
	PostgresDBName   string      `env:"POSTGRES_DBNAME,required"`
	PostgresAPPName  string      `env:"POSTGRES_APPNAME,required"`
	PostgresSSLMode  string      `env:"POSTGRES_SSLMODE" envDefault:"disable"`
}

func (Configs) DB() string {
	return "postgres"
}

func (c Configs) DSN() string {
	return fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v application_name=%v sslmode=%v",
		strings.Join(c.PostgresHosts, ","),
		c.PostgresPort,
		c.PostgresDBName,
		c.PostgresUsername,
		c.PostgresPassword,
		c.PostgresAPPName,
		c.PostgresSSLMode,
	)
}
