package muxext

import "github.com/gorilla/mux"

type ControllerRegister interface {
	RegisterController(router *mux.Router)
}
