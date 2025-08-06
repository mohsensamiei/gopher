package mediator

import (
	"context"
	"fmt"
)

type Request any
type Response any

type CommandHandler[TRequest Request, TResponse Response] interface {
	Handle(ctx context.Context, request TRequest) (TResponse, error)
}

func RegisterCommandHandlerOn[TRequest Request, TResponse Response](m *Mediator, h CommandHandler[TRequest, TResponse]) {
	key := getTypeKey[TRequest]()
	m.requests.Store(key, h)
}

func SendOn[TRequest Request, TResponse Response](m *Mediator, ctx context.Context, req TRequest) (TResponse, error) {
	var handler CommandHandler[TRequest, TResponse]
	{
		key := getTypeKey[TRequest]()
		handlerRaw, ok := m.requests.Load(key)
		if !ok {
			var zero TResponse
			return zero, fmt.Errorf("handler not found for request %v", key)
		}
		handler = handlerRaw.(CommandHandler[TRequest, TResponse])
	}
	chain := func(ctx context.Context, req TRequest) (TResponse, error) {
		return handler.Handle(ctx, req)
	}
	m.behaviorMutex.RLock()
	defer m.behaviorMutex.RUnlock()
	for i := len(m.behaviors) - 1; i >= 0; i-- {
		bRaw := m.behaviors[i]

		if b, ok := bRaw.(PipelineBehavior[TRequest, TResponse]); ok {
			next := chain
			chain = func(ctx context.Context, req TRequest) (TResponse, error) {
				return b.Handle(ctx, req, next)
			}
		}
	}
	return chain(ctx, req)
}
