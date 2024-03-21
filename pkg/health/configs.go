package health

import "github.com/mohsensamiei/gopher/v2/pkg/netext"

type Configs struct {
	HealthPort netext.Port `env:"HEALTH_PORT" envDefault:"5000"`
}
