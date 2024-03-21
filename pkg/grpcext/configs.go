package grpcext

import "github.com/mohsensamiei/gopher/v2/pkg/netext"

type Configs struct {
	GrpcPort netext.Port `env:"GRPC_PORT" envDefault:"6565"`
}
