package main

import (
	"context"
	"github.com/pinosell/gopher/pkg/gormext"
	"github.com/pinosell/gopher/pkg/postgresext"
	"github.com/pinosell/gopher/pkg/query"
	log "github.com/sirupsen/logrus"
)

type User struct {
	gormext.IncrementalModel
}

func (u User) FullTextName() string {
	return "FullName"
}

type UserRepository struct {
	gormext.CrudRepository[User, uint32]
}

func main() {
	r := new(UserRepository)
	log.Print(r.ReturnByID(context.Background(), 123, query.Empty))

	configs := &postgresext.Configs{
		PostgresHosts:    []string{"localhost"},
		PostgresPort:     5432,
		PostgresUsername: "postgres",
		PostgresPassword: "postgres",
		PostgresDBName:   "example",
		PostgresAPPName:  "goraz",
		PostgresSSLMode:  "disable",
	}
	db, err := gormext.Connect(configs)
	if err != nil {
		panic(err)
	}
	list := new([]*User)
	gormext.ApplyQuery[User](db, query.Search("text")).Find(&list)
	log.Print(list)
}
