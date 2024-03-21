package worker

import "github.com/mohsensamiei/gopher/pkg/netext"

type Configs struct {
	WorkerPort netext.Port `env:"WORKER_PORT" envDefault:"7337"`
}
