package gormext

import (
	"github.com/pinosell/gopher/pkg/errors"
	"google.golang.org/grpc/codes"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection interface {
	DB() string
	DSN() string
}

var (
	connections = map[string]func(dsn string) (*gorm.DB, error){
		"postgres": connectPostgres,
		"mysql":    connectMysql,
	}
)

func Connect(conn Connection) (*gorm.DB, error) {
	f, ok := connections[conn.DB()]
	if !ok {
		return nil, errors.New(codes.Unimplemented).
			WithDetailF("connection does not support '%v'", conn.DB())
	}
	return f(conn.DSN())
}

func connectPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		FullSaveAssociations: false,
		Logger:               NewLogger(),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectMysql(dns string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dns), &gorm.Config{
		FullSaveAssociations: false,
		Logger:               NewLogger(),
	})
}
