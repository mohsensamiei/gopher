package worker

import (
	"context"
	"github.com/mohsensamiei/gopher/pkg/service"
	"net"
	"net/http"
)

func Serve(configs Configs, workers []Worker, builders []Builder) {
	ctx, cancel := context.WithCancel(context.Background())
	service.Serve(configs.WorkerPort, Platform, func(lst net.Listener) error {
		defer cancel()
		for _, builder := range builders {
			ctx = builder(ctx)
		}
		for _, w := range workers {
			go func(ctx context.Context, worker Worker) {
				worker.RegisterWorker(ctx)
			}(ctx, w)
		}
		return http.Serve(lst, new(Handler))
	})
}
