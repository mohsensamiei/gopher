package main

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/gormext"
	"github.com/mohsensamiei/gopher/v2/pkg/postgresext"
	"github.com/mohsensamiei/gopher/v2/pkg/query"
	log "github.com/sirupsen/logrus"
)

type User struct {
	gormext.IncrementalModel
}

func (u User) FullTextName() string {
	return "FullName"
}

type UserRepository struct {
	gormext.CrudRepository[User]
}

func main() {
	r := new(UserRepository)
	log.Print(r.ReturnByPK(context.Background(), query.Empty, 123))

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
