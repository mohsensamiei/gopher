package worker

import "github.com/mohsensamiei/gopher/v2/pkg/netext"

type Configs struct {
	WorkerPort netext.Port `env:"WORKER_PORT" envDefault:"7337"`
}
