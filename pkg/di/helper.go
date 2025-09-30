package di

import (
	"reflect"
)

func TypeName[T any]() string {
	var x T
	t := reflect.TypeOf(x)
	if t == nil {
		t = reflect.TypeOf((*T)(nil)).Elem()
	}
	if t.PkgPath() == "" {
		return t.Name()
	}
	return t.PkgPath() + "." + t.Name()
}
