package muxext

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

type key[T comparable] interface {
	SetID(id T)
	ParseID(id string) (T, error)
	SetSlug(slug string)
}

func PathKey[T comparable](req *http.Request, name string, out key[T]) {
	v := mux.Vars(req)[name]
	id, err := out.ParseID(v)
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
