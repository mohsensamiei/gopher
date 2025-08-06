package grpcext

import (
	"github.com/mohsensamiei/gopher/v3/pkg/service"
	"google.golang.org/grpc"
	"net"
)

func Serve(configs Configs, services []ServiceRegister, options []grpc.ServerOption) {
	service.Serve(configs.GrpcPort, "grpc", func(lst net.Listener) error {
		server := grpc.NewServer(options...)
		for _, srv := range services {
			srv.RegisterService(server)
		}
		return server.Serve(lst)
	})
}
