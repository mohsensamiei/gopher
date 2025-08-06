package health

import "github.com/mohsensamiei/gopher/v3/pkg/netext"

type Configs struct {
	HealthPort netext.Port `env:"HEALTH_PORT" envDefault:"5000"`
}
