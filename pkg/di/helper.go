package di

import (
	"fmt"
)

func TypeName[T any]() string {
	var t T
	return fmt.Sprintf("%T", t)
}
