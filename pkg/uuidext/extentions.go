package uuidext

import (
	"github.com/google/uuid"
	"github.com/pinosell/gopher/pkg/errors"
	"google.golang.org/grpc/codes"
)

func Parse(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, codes.InvalidArgument)
	}
	return id, nil
}

func ParseSlice(list []string) ([]uuid.UUID, error) {
	var result []uuid.UUID
	for _, item := range list {
		id, err := Parse(item)
		if err != nil {
			return nil, err
		}
		result = append(result, id)
	}
	return result, nil
}
