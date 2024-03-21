package httpext

import (
	"encoding/json"
	"net/http"
)

func BindRequestModel(req *http.Request, model any) error {
	payload, err := RequestPayload(req)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(payload, model); err != nil {
		return err
	}
	return nil
}

func BindResponseModel(res *http.Response, model any) ([]byte, error) {
	payload, err := ResponsePayload(res)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(payload, model); err != nil {
		return nil, err
	}
	return payload, nil
}
