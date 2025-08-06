package migrateext

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"google.golang.org/grpc/codes"
)

type Connection interface {
	DB() string
	DSN() string
}

func open(conn Connection) (database.Driver, error) {
	db, err := sql.Open(conn.DB(), conn.DSN())
	if err != nil {
		return nil, err
	}

	var driver database.Driver
	switch conn.DB() {
	case "postgres":
		driver, err = postgres.WithInstance(db, &postgres.Config{
			MultiStatementEnabled: true,
		})
	default:
		return nil, errors.New(codes.Unimplemented)
	}
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func Up(conn Connection, path string) error {
	driver, err := open(conn)
	if err != nil {
		return err
	}

	var migrator *migrate.Migrate
	migrator, err = migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%v", path), conn.DB(), driver)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = migrator.Close()
	}()

	if err = migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func Down(conn Connection, path string) error {
	driver, err := open(conn)
	if err != nil {
		return err
	}

	var migrator *migrate.Migrate
	migrator, err = migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%v", path), conn.DB(), driver)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = migrator.Close()
	}()

	if err = migrator.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
