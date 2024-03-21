package muxext

import (
	"github.com/gorilla/mux"
	"github.com/mohsensamiei/gopher/pkg/httpext"
	"github.com/mohsensamiei/gopher/pkg/service"
	"github.com/rs/cors"
	"net"
	"net/http"
)

func Serve(configs httpext.Configs, controllers []ControllerRegister, middlewares []mux.MiddlewareFunc, cors *cors.Cors) {
	service.Serve(configs.HttpPort, httpext.Platform, func(lst net.Listener) error {
		router := NewRouter()
		router.Use(middlewares...)
		for _, ctrl := range controllers {
			ctrl.RegisterController(router)
		}
		return http.Serve(lst, cors.Handler(router))
	})
}
