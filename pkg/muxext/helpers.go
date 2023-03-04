package muxext

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

type key interface {
	SetID(id uuid.UUID)
	SetSlug(slug string)
}

func PathKey(req *http.Request, name string, out key) {
	v := mux.Vars(req)[name]
	id, err := uuid.Parse(v)
	if err == nil {
		out.SetID(id)
		return
	}
	out.SetSlug(v)
}

func PathParam(req *http.Request, name string) string {
	return mux.Vars(req)[name]
}

func HandleFunc(router *mux.Router, path string, f http.HandlerFunc, middlewares ...mux.MiddlewareFunc) *mux.Route {
	handler := http.Handler(f)
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return router.Handle(path, handler)
}
