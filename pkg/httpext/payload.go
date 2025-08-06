package httpext

import (
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"google.golang.org/grpc/codes"
	"io"
	"net/http"
)

func RequestPayload(req *http.Request) ([]byte, error) {
	defer func() {
		_ = req.Body.Close()
	}()
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, codes.InvalidArgument)
	}
	return payload, nil
}

func ResponsePayload(res *http.Response) ([]byte, error) {
	defer func() {
		_ = res.Body.Close()
	}()
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, codes.InvalidArgument)
	}
	return payload, nil
}
