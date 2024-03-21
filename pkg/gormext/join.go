package gormext

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/mohsensamiei/gopher/pkg/slices"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func join[T any](db *gorm.DB, query string) *gorm.DB {
	tai := new(T)
	tav := reflect.ValueOf(tai)

	fk, ok := tav.Elem().Type().FieldByName(query)
	if !ok {
		return db
	}

	isSlice := false
	var tbv reflect.Value
	{
		fkt := fk.Type
		if fkt.Kind() == reflect.Ptr {
			fkt = fkt.Elem()
		}
		switch fkt.Kind() {
		case reflect.Slice:
			isSlice = true
			fkt = fkt.Elem()
			if fkt.Kind() == reflect.Ptr {
				fkt = fkt.Elem()
			}
			tbv = reflect.New(fkt)
		case reflect.Struct:
			tbv = reflect.New(fkt)
		default:
			return db
		}
	}
	ta := TableName(db, tav.Interface())
	tb := TableName(db, tbv.Interface())
	tba := strcase.ToSnake(query)

	var joins []string
	for _, j := range db.Statement.Joins {
		var t, a string
		_, _ = fmt.Sscanf(j.Name, "join %s %s on", &t, &a)
		joins = append(joins, a)
	}
	if slices.Contains(tba, joins...) {
		return db
	}

	tag := make(map[string]string)
	for _, j := range strings.Split(fk.Tag.Get("gorm"), ";") {
		d := strings.Split(j, ":")
		if len(d) != 2 {
			continue
		}
		tag[strings.ToLower(d[0])] = d[1]
	}

	if isSlice {
		db = db.Joins(fmt.Sprintf("join %v %v on %v.%v = %v.%v", tb, tba,
			tba, strcase.ToSnake(tag["foreignkey"]),
			ta, strcase.ToSnake(tag["references"]),
		))
	} else {
		db = db.Joins(fmt.Sprintf("join %v %v on %v.%v = %v.%v", tb, tba,
			tba, strcase.ToSnake(tag["references"]),
			ta, strcase.ToSnake(tag["foreignkey"]),
		))
	}
	if !db.Statement.Unscoped {
		for i := 0; i < tbv.Elem().NumField(); i++ {
			fv := tbv.Elem().Field(i)
			switch fv.Interface().(type) {
			case gorm.DeletedAt:
				ft := tbv.Elem().Type().Field(i)
				db = db.Where(fmt.Sprintf("%v.%v IS NULL", tba, strcase.ToSnake(ft.Name)))
			}
		}
	}
	return db
}
