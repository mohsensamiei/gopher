package templates

const (
	CmdMain = `
package main

import (
	// GOPHER: Don't remove this line
	// {{ .import }}
	"github.com/mohsensamiei/gopher/v2/pkg/envext"
	"github.com/mohsensamiei/gopher/v2/pkg/grpcext"
	"github.com/mohsensamiei/gopher/v2/pkg/health"
	"github.com/mohsensamiei/gopher/v2/pkg/httpext"
	"github.com/mohsensamiei/gopher/v2/pkg/i18next"
	"github.com/mohsensamiei/gopher/v2/pkg/logext"
	"github.com/mohsensamiei/gopher/v2/pkg/muxext"
	"github.com/mohsensamiei/gopher/v2/pkg/service"
	"github.com/mohsensamiei/gopher/v2/pkg/closer"
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

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Format: Bearer [Access Token]
func main() {
	defer closer.Defer()

	health.Serve(configs.Health)

	grpcext.Serve(configs.Grpc, []grpcext.ServiceRegister{
		// GOPHER: Don't remove this line
		// {{ .service }}
	}, []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpcext.UnaryWrapErrorInterceptor(),
			grpcext.UnaryContextMetadataInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
	})

	muxext.Serve(Service, configs.Http, []muxext.ControllerRegister{
		// GOPHER: Don't remove this line
		// {{ .controller }}
	}, []mux.MiddlewareFunc{
		muxext.LangMiddleware(),
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
