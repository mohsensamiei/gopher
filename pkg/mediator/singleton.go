package mediator

import (
	"context"
)

var defaultMediator = New()

func UseBehavior[TRequest Request, TResponse Response](b PipelineBehavior[TRequest, TResponse]) {
	UseBehaviorOn[TRequest, TResponse](defaultMediator, b)
}

func RegisterNotificationHandler[T Notification](h NotificationHandler[T]) {
	RegisterNotificationHandlerOn[T](defaultMediator, h)
}

func Publish[T Notification](ctx context.Context, n T) error {
	return PublishOn[T](defaultMediator, ctx, n)
}

func RegisterCommandHandler[TRequest Request, TResponse Response](h CommandHandler[TRequest, TResponse]) {
	RegisterCommandHandlerOn[TRequest, TResponse](defaultMediator, h)
}

func Send[TRequest Request, TResponse any](ctx context.Context, req TRequest) (TResponse, error) {
	return SendOn[TRequest, TResponse](defaultMediator, ctx, req)
}
