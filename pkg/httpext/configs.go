package httpext

import "github.com/pinosell/gopher/pkg/netext"

type Configs struct {
	HttpPort netext.Port `env:"HTTP_PORT" envDefault:"8080"`
}
