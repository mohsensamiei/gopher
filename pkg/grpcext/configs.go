package grpcext

import "github.com/mohsensamiei/gopher/pkg/netext"

type Configs struct {
	GrpcPort netext.Port `env:"GRPC_PORT" envDefault:"6565"`
}
