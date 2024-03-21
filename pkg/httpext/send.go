package httpext

import (
	"encoding/json"
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/mohsensamiei/gopher/v2/pkg/i18next"
	"github.com/mohsensamiei/gopher/v2/pkg/mimeext"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SendCode(res http.ResponseWriter, _ *http.Request, code int) {
	res.WriteHeader(code)
}

func SendError(res http.ResponseWriter, req *http.Request, err error) {
	model := errors.Cast(err)
	if model.Localize() == "" {
		model = model.SetLocalize(i18next.ByContext(req.Context(), model.Slug()))
	}
	SendModel(res, req, model.StatusCode(), model)
}

func SendModel(res http.ResponseWriter, req *http.Request, code int, model any) {
	bytes, err := json.Marshal(model)
	if err != nil {
		log.WithError(err).Error("can not marshal model")
	}
	Send(res, req, code, mimeext.Json, bytes)
}

func Send(res http.ResponseWriter, _ *http.Request, code int, mime string, data []byte) {
	res.Header().Set(ContentTypeHeader, mime)
	res.Header().Set(CharsetHeader, "utf-8")
	res.WriteHeader(code)
	if _, err := res.Write(data); err != nil {
		log.WithError(err).Error("can not write data on response")
	}
}
