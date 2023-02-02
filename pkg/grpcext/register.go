package grpcext

import "google.golang.org/grpc"

type ServiceRegister interface {
	RegisterService(server *grpc.Server)
}
