package main

import (
	"github.com/pinosell/gopher/pkg/migrateext"
	"github.com/pinosell/gopher/pkg/postgresext"
)

func main() {
	configs := &postgresext.Configs{
		PostgresHosts:    []string{"localhost"},
		PostgresPort:     5432,
		PostgresUsername: "postgres",
		PostgresPassword: "postgres",
		PostgresDBName:   "example",
		PostgresAPPName:  "goraz",
		PostgresSSLMode:  "disable",
	}
	path := "./pkg/migrateext/examples/migrations"

	if err := migrateext.Up(configs, path); err != nil {
		panic(err)
	}
	if err := migrateext.Down(configs, path); err != nil {
		panic(err)
	}
}
