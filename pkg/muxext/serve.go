package muxext

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mohsensamiei/gopher/v2/pkg/httpext"
	"github.com/mohsensamiei/gopher/v2/pkg/service"
	"github.com/rs/cors"
	"net"
	"net/http"
)

func Serve(serviceName string, configs httpext.Configs, controllers []ControllerRegister, middlewares []mux.MiddlewareFunc, cors *cors.Cors) {
	service.Serve(configs.HttpPort, httpext.Platform, func(lst net.Listener) error {
		router := NewRouter(fmt.Sprintf("/%v", serviceName))
		router.Use(middlewares...)
		for _, ctrl := range controllers {
			ctrl.RegisterController(router)
		}
		return http.Serve(lst, cors.Handler(router))
	})
}
