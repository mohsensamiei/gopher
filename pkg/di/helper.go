package di

import (
	"fmt"
	"reflect"
)

func TypeName[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()
	return typeString(t)
}

func typeString(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Ptr:
		return "*" + typeString(t.Elem())
	case reflect.Slice:
		return "[]" + typeString(t.Elem())
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", t.Len(), typeString(t.Elem()))
	case reflect.Map:
		return "map[" + typeString(t.Key()) + "]" + typeString(t.Elem())
	case reflect.Chan:
		return "chan " + typeString(t.Elem())
	case reflect.Func:
		return "func"
	default:
		if t.PkgPath() == "" {
			return t.Name()
		}
		return t.PkgPath() + "." + t.Name()
	}
}
