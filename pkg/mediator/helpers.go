package mediator

import (
	"fmt"
)

func getTypeKey[T any]() string {
	var t T
	return fmt.Sprintf("%T", t)
}
