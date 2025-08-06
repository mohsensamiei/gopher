package mediator

import (
	"context"
	"sync"
)

type Notification any

type NotificationHandler[T Notification] interface {
	Handle(ctx context.Context, n T) error
}

func RegisterNotificationHandlerOn[T Notification](m *Mediator, h NotificationHandler[T]) {
	key := getTypeKey[T]()
	raw, _ := m.notifications.LoadOrStore(key, &[]NotificationHandler[T]{})
	handlers := raw.(*[]NotificationHandler[T])
	*handlers = append(*handlers, h)
}

func PublishOn[T Notification](m *Mediator, ctx context.Context, n T) error {
	hRaw, ok := m.notifications.Load(getTypeKey[T]())
	if !ok {
		return nil
	}
	handlers := *(hRaw.(*[]NotificationHandler[T]))

	wg := sync.WaitGroup{}
	errCh := make(chan error, len(handlers))
	for _, h := range handlers {
		h := h
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Handle(ctx, n); err != nil {
				errCh <- err
			}
		}()
	}
	wg.Wait()
	close(errCh)

	for err := range errCh {
		return err
	}
	return nil
}
