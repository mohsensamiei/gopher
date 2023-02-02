package gormext

import (
	"gorm.io/gorm"
	"reflect"
)

type Tabler interface {
	TableName() string
}

func TableName(db *gorm.DB, model any) string {
	if tabler, ok := model.(Tabler); ok {
		return tabler.TableName()
	}
	return db.NamingStrategy.TableName(reflect.TypeOf(model).Elem().Name())
}
