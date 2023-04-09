package worker

import (
	"context"
	"github.com/pinosell/gopher/pkg/service"
	"net"
	"net/http"
)

type DI map[string]func(ctx context.Context) any

func Serve(configs Configs, workers []Worker, di DI) {
	ctx, cancel := context.WithCancel(context.Background())
	for k, f := range di {
		ctx = context.WithValue(ctx, k, f(ctx))
	}
	service.Serve(configs.WorkerPort, Platform, func(lst net.Listener) error {
		defer cancel()
		for _, w := range workers {
			go func(ctx context.Context, worker Worker) {
				worker.RegisterWorker(ctx)
			}(ctx, w)
		}
		return http.Serve(lst, new(Handler))
	})
}
