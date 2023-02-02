package grpcext

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func UnaryToServerOption(mdl ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		mdl...,
	))
}

func StreamToServerOption(mdl ...grpc.StreamServerInterceptor) grpc.ServerOption {
	return grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		mdl...,
	))
}
