package health

import (
	"github.com/mohsensamiei/gopher/v2/pkg/service"
	"net"
	"net/http"
)

func Serve(configs Configs) {
	service.Serve(configs.HealthPort, Platform, func(lst net.Listener) error {
		return http.Serve(lst, NewHandler())
	})
}
