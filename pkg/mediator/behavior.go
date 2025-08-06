package mediator

import "context"

type PipelineBehavior[TRequest any, TResponse any] interface {
	Handle(ctx context.Context, req TRequest, next func(context.Context, TRequest) (TResponse, error)) (TResponse, error)
}

func UseBehaviorOn[TRequest any, TResponse any](m *Mediator, b PipelineBehavior[TRequest, TResponse]) {
	m.behaviorMutex.Lock()
	defer m.behaviorMutex.Unlock()
	m.behaviors = append(m.behaviors, b)
}
