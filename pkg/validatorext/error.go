package validatorext

import (
	"fmt"
	"github.com/mohsensamiei/gopher/pkg/errors"
	"google.golang.org/grpc/codes"
)

type Validation struct {
	Field string
	Tag   string
}

func ParseError(err error) []*Validation {
	st := errors.Cast(err)
	if st == nil || st.Code() != codes.InvalidArgument {
		return nil
	}
	var res []*Validation
	for _, detail := range errors.Cast(err).Details() {
		v := new(Validation)
		if _, err = fmt.Sscanf(detail, "%s %s", &v.Field, &v.Tag); err == nil {
			res = append(res, v)
		}
	}
	return res
}
