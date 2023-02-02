package templates

const (
	AppController = `
package {{ .plural }}

import (
	"{{ .repository }}/api"
	"github.com/pinosell/gopher/pkg/grpcext"
	"github.com/pinosell/gopher/pkg/httpext"
	"github.com/pinosell/gopher/pkg/muxext"
	"github.com/gorilla/mux"
	"net/http"
)

func NewController(configs grpcext.Configs) muxext.ControllerRegister {
	internalConn := grpcext.NewInternalConnection(configs)
	return &Controller{
		{{ .Singular }}ServiceClient: api.New{{ .Singular }}ServiceClient(internalConn),
	}
}

type Controller struct {
	api.{{ .Singular }}ServiceClient
}

func (c Controller) RegisterController(router *mux.Router) {
	muxext.HandleFunc(router, "/{{ .plural }}", c.List).Methods(http.MethodGet)
	muxext.HandleFunc(router, "/{{ .plural }}", c.Create).Methods(http.MethodPost)
	muxext.HandleFunc(router, "/{{ .plural }}/{{ "{" }}{{ .singular }}_id}", c.Return).Methods(http.MethodGet)
	muxext.HandleFunc(router, "/{{ .plural }}/{{ "{" }}{{ .singular }}_id}", c.Update).Methods(http.MethodPut)
	muxext.HandleFunc(router, "/{{ .plural }}/{{ "{" }}{{ .singular }}_id}", c.Delete).Methods(http.MethodDelete)
}

func (c Controller) List(res http.ResponseWriter, req *http.Request) {
	model := &api.{{ .Singular }}List{
		Query: req.URL.RawQuery,
	}
	result, err := c.{{ .Singular }}ServiceClient.List(req.Context(), model)
	if err != nil {
		httpext.SendError(res, req, err)
		return
	}
	httpext.SendModel(res, req, http.StatusOK, result)
}

func (c Controller) Create(res http.ResponseWriter, req *http.Request) {
	model := &api.{{ .Singular }}Create{
		Query: req.URL.RawQuery,
	}
	if err := httpext.BindModel(req, model); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	result, err := c.{{ .Singular }}ServiceClient.Create(req.Context(), model)
	if err != nil {
		httpext.SendError(res, req, err)
		return
	}
	httpext.SendModel(res, req, http.StatusCreated, result)
}

func (c Controller) Return(res http.ResponseWriter, req *http.Request) {
	model := &api.{{ .Singular }}Return{
		{{ .Singular }}ID:    muxext.PathParam(req, "{{ .singular }}_id"),
		Query: req.URL.RawQuery,
	}
	result, err := c.{{ .Singular }}ServiceClient.Return(req.Context(), model)
	if err != nil {
		httpext.SendError(res, req, err)
		return
	}
	httpext.SendModel(res, req, http.StatusOK, result)
}

func (c Controller) Update(res http.ResponseWriter, req *http.Request) {
	model := &api.{{ .Singular }}Update{
		{{ .Singular }}ID:    muxext.PathParam(req, "{{ .singular }}_id"),
		Query: req.URL.RawQuery,
	}
	if err := httpext.BindModel(req, model); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	result, err := c.{{ .Singular }}ServiceClient.Update(req.Context(), model)
	if err != nil {
		httpext.SendError(res, req, err)
		return
	}
	httpext.SendModel(res, req, http.StatusOK, result)
}

func (c Controller) Delete(res http.ResponseWriter, req *http.Request) {
	model := &api.{{ .Singular }}Delete{
		{{ .Singular }}ID: muxext.PathParam(req, "{{ .singular }}_id"),
	}
	if _, err := c.{{ .Singular }}ServiceClient.Delete(req.Context(), model); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	httpext.SendCode(res, req, http.StatusNoContent)
}
`

	AppService = `
package {{ .plural }}

import (
	"context"
	"{{ .repository }}/api"
	"github.com/pinosell/gopher/pkg/errors"
	"github.com/pinosell/gopher/pkg/grpcext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func NewService() grpcext.ServiceRegister {
	return new(Service)
}

type Service struct {
}

func (s Service) RegisterService(server *grpc.Server) {
	api.Register{{ .Singular }}ServiceServer(server, s)
}

func (Service) List(ctx context.Context, req *api.{{ .Singular }}List) (*api.{{ .Plural }}, error) {
	return nil, errors.New(codes.Unimplemented)
}

func (s Service) Create(ctx context.Context, req *api.{{ .Singular }}Create) (*api.{{ .Singular }}, error) {
	return nil, errors.New(codes.Unimplemented)
}

func (s Service) Return(ctx context.Context, req *api.{{ .Singular }}Return) (*api.{{ .Singular }}, error) {
	return nil, errors.New(codes.Unimplemented)
}

func (s Service) Update(ctx context.Context, req *api.{{ .Singular }}Update) (*api.{{ .Singular }}, error) {
	return nil, errors.New(codes.Unimplemented)
}

func (Service) Delete(ctx context.Context, req *api.{{ .Singular }}Delete) (*api.Void, error) {
	return nil, errors.New(codes.Unimplemented)
}
`
)
