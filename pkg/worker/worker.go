package worker

import "context"

type Worker interface {
	RegisterWorker(ctx context.Context)
}
