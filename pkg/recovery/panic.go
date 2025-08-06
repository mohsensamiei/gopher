package recovery

import "github.com/pkg/errors"

func handlePanic(r interface{}) error {
	var errWithStack error
	if err, ok := r.(error); ok {
		errWithStack = errors.WithStack(err)
	} else {
		errWithStack = errors.Errorf("%+v", r)
	}
	return errWithStack
}
