package muxext

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
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
