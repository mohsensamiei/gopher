package health

import (
	"github.com/mohsensamiei/gopher/v3/pkg/service"
	"net"
	"net/http"
)

func Serve(configs Configs) {
	service.Serve(configs.HealthPort, "health", func(lst net.Listener) error {
		return http.Serve(lst, NewHandler())
	})
}
