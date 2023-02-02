package httpext

import (
	"bytes"
	"encoding/json"
	"github.com/pinosell/gopher/pkg/errors"
	"google.golang.org/grpc/codes"
	"net/http"
)

func BindModel(req *http.Request, model any) (err error) {
	defer func() {
		_ = req.Body.Close()
		if err != nil {
			err = errors.Wrap(err, codes.InvalidArgument)
		}
	}()
	var obj map[string]any
	if err = json.NewDecoder(req.Body).Decode(&obj); err != nil {
		return err
	}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err = encoder.Encode(obj); err != nil {
		return err
	}
	return json.NewDecoder(&buf).Decode(model)
}
