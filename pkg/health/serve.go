package health

import (
	"github.com/pinosell/gopher/pkg/service"
	"net"
	"net/http"
)

func Serve(configs Configs) {
	service.Serve(configs.HealthPort, Platform, func(lst net.Listener) error {
		return http.Serve(lst, NewHandler())
	})
}
