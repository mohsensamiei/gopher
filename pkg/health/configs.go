package health

import "github.com/pinosell/gopher/pkg/netext"

type Configs struct {
	HealthPort netext.Port `env:"HEALTH_PORT" envDefault:"5000"`
}
