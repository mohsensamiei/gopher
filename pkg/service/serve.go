package service

import (
	"fmt"
	"github.com/mohsensamiei/gopher/v2/pkg/netext"
	"net"
)

var (
	serves = make(map[netext.Port]*serve)
)

func Serve(port netext.Port, name string, f ServeFunc) {
	serves[port] = &serve{
		Name:     name,
		Function: f,
	}
}

type ServeFunc func(listener net.Listener) error

type serve struct {
	Name     string
	Function ServeFunc
}

func (s serve) Listen(port netext.Port) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}
	defer func() {
		_ = listener.Close()
	}()
	return s.Function(listener)
}
