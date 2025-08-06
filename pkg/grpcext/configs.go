package grpcext

import "github.com/mohsensamiei/gopher/v3/pkg/netext"

type Configs struct {
	GrpcPort netext.Port `env:"GRPC_PORT" envDefault:"6565"`
}
