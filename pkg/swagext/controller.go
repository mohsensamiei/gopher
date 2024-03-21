package swagext

import (
	"github.com/gorilla/mux"
	"github.com/mohsensamiei/gopher/v2/pkg/httpext"
	"github.com/mohsensamiei/gopher/v2/pkg/mimeext"
	"github.com/mohsensamiei/gopher/v2/pkg/muxext"
	"github.com/swaggo/swag"
	"net/http"
	"net/url"
)

func NewMuxController(configs Configs) *MuxController {
	return &MuxController{
		Configs: configs,
	}
}

type MuxController struct {
	Configs
}

func (c MuxController) RegisterController(router *mux.Router) {
	muxext.HandleFunc(router, "/swagger", c.Swagger).Methods(http.MethodGet)
}

func (c MuxController) Swagger(res http.ResponseWriter, req *http.Request) {
	doc := swag.GetSwagger(Name).(*swag.Spec)
	{
		uri, err := url.Parse(c.ExternalURL)
		if err != nil {
			httpext.SendError(res, req, err)
			return
		}
		doc.Host = uri.Host
		doc.Schemes = []string{uri.Scheme}
	}
	httpext.Send(res, req, http.StatusOK, mimeext.Json, []byte(doc.ReadDoc()))
}
