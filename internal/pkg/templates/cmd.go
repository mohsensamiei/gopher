package templates

const (
	CmdMain = `
package main

import (
	// GOPHER: Don't remove this line
	// {{ .import }}
	"github.com/pinosell/gopher/pkg/envext"
	"github.com/pinosell/gopher/pkg/grpcext"
	"github.com/pinosell/gopher/pkg/health"
	"github.com/pinosell/gopher/pkg/httpext"
	"github.com/pinosell/gopher/pkg/i18next"
	"github.com/pinosell/gopher/pkg/logext"
	"github.com/pinosell/gopher/pkg/metadataext"
	"github.com/pinosell/gopher/pkg/muxext"
	"github.com/pinosell/gopher/pkg/service"
	"github.com/gorilla/mux"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Configs struct {
	Health   health.Configs
	Grpc     grpcext.Configs
	Http     httpext.Configs
	Log      logext.Configs
	I18N     i18next.Configs
}

const (
	Service = "{{ .name }}"
)

var (
	Version = "NAN"
	configs Configs
)

func init() {
	logext.Initial(Service, Version)
	if err := envext.Parse(&configs); err != nil {
		log.WithError(err).Panic("can not parse env configs")
	}
	logext.Setup(Service, Version, configs.Log)

	if err := i18next.Setup(configs.I18N, "assets/locales"); err != nil {
		log.WithError(err).Panic("can not setup i18n package")
	}
}

func main() {
	health.Serve(configs.Health)

	grpcext.Serve(configs.Grpc, []grpcext.ServiceRegister{
		// GOPHER: Don't remove this line
		// {{ .service }}
	}, []grpc.ServerOption{
		grpcext.UnaryToServerOption(
			grpcext.UnaryWrapErrorInterceptor(),
			metadataext.UnaryContextInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
	})

	muxext.Serve(configs.Http, []muxext.ControllerRegister{
		// GOPHER: Don't remove this line
		// {{ .controller }}
	}, []mux.MiddlewareFunc{
	}, httpext.DefaultCors)

	service.Start()
}
`
	CmdImport = `
	// {{ .import }}
	"{{ .repository }}/internal/app/{{ .plural }}"
`
	CmdService = `
		// {{ .service }}
		{{ .plural }}.NewService(),
`
	CmdController = `
		// {{ .controller }}
		{{ .plural }}.NewController(configs.Grpc),
`
)
