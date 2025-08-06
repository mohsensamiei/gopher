package worker

import "github.com/mohsensamiei/gopher/v3/pkg/netext"

type Configs struct {
	WorkerPort netext.Port `env:"WORKER_PORT" envDefault:"7337"`
}
