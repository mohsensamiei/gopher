package httpext

import "github.com/mohsensamiei/gopher/pkg/netext"

type Configs struct {
	HttpPort netext.Port `env:"HTTP_PORT" envDefault:"8080"`
}
