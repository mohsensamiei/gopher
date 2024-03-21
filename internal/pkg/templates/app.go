package templates

const (
	AppController = `
package {{ .plural }}

import (
	"{{ .repository }}/api"
	"github.com/mohsensamiei/gopher/v2/pkg/grpcext"
	"github.com/mohsensamiei/gopher/v2/pkg/httpext"
	"github.com/mohsensamiei/gopher/v2/pkg/muxext"
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

//	@Summary	List of {{ .plural }}
//	@Tags		{{ .plural }}
//	@Router		/{{ .command }}/{{ .plural }} [get]
//	@Param		{object}	query	queryext.Multiple	false	"Query string"
//	@Produce	json
//	@Success	200	{object}	api.{{ .Plural }}
//	@Failure	400	{object}	errors.Model
//	@Failure	404	{object}	errors.Model
//	@Failure	500	{object}	errors.Model
//	@Failure	501	{object}	errors.Model
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

//	@Summary	Creates a {{ .singular }}
//	@Tags		{{ .plural }}
//	@Router		/{{ .command }}/{{ .plural }} [post]
//	@Security	BearerAuth
//	@Param		{object}	query	queryext.Single	false	"Query string"
//	@Accept		json
//	@Param		{object}	body		api.{{ .Singular }}Create	true	"Request body"
//	@Produce	json
//	@Success	201			{object}	api.{{ .Singular }}
//	@Failure	400			{object}	errors.Model
//	@Failure	401			{object}	errors.Model
//	@Failure	403			{object}	errors.Model
//	@Failure	404			{object}	errors.Model
//	@Failure	500			{object}	errors.Model
//	@Failure	501			{object}	errors.Model
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

//	@Summary	Returns a {{ .singular }}
//	@Tags		{{ .plural }}
//	@Router		/{{ .command }}/{{ .plural }}/{{{ .singular }}_id} [get]
//	@Param		{object}	query	queryext.Single	false	"Query string"
//	@Param		{{ .singular }}_id	path	string			true	"{{ .Singular }} primary key"
//	@Produce	json
//	@Success	200	{object}	api.{{ .Singular }}
//	@Failure	400	{object}	errors.Model
//	@Failure	404	{object}	errors.Model
//	@Failure	500	{object}	errors.Model
//	@Failure	501	{object}	errors.Model
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

//	@Summary	Updates a {{ .singular }}
//	@Tags		{{ .plural }}
//	@Router		/{{ .command }}/{{ .plural }}/{{{ .singular }}_id} [put]
//	@Security	BearerAuth
//	@Param		{object}	query	queryext.Single	false	"Query string"
//	@Param		{{ .singular }}_id	path	string			true	"{{ .Singular }} primary key"
//	@Produce	json
//	@Param		{object}	body		api.{{ .Singular }}Update	true	"Request body"
//	@Accept		json
//	@Success	200			{object}	api.{{ .Singular }}
//	@Failure	400			{object}	errors.Model
//	@Failure	401			{object}	errors.Model
//	@Failure	403			{object}	errors.Model
//	@Failure	404			{object}	errors.Model
//	@Failure	500			{object}	errors.Model
//	@Failure	501			{object}	errors.Model
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

//	@Summary	Deletes a {{ .singular }}
//	@Tags		{{ .plural }}
//	@Router		/{{ .command }}/{{ .plural }}/{{{ .singular }}_id} [delete]
//	@Security	BearerAuth
//	@Param		{{ .singular }}_id	path	string	true	"{{ .Singular }} primary key"
//	@Success	204
//	@Failure	400	{object}	errors.Model
//	@Failure	401	{object}	errors.Model
//	@Failure	403	{object}	errors.Model
//	@Failure	404	{object}	errors.Model
//	@Failure	500	{object}	errors.Model
//	@Failure	501	{object}	errors.Model
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
	"github.com/mohsensamiei/gopher/v2/pkg/errors"
	"github.com/mohsensamiei/gopher/v2/pkg/grpcext"
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
