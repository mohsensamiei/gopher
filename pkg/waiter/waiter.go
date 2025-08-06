package waiter

import (
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"google.golang.org/grpc/codes"
	"sync"
	"time"
)

type Waiter struct {
	mu     sync.Mutex
	waitCh map[string]chan struct{}
}

func NewWaiter() *Waiter {
	return &Waiter{
		waitCh: make(map[string]chan struct{}),
	}
}

func (w *Waiter) WaitForWithTimeout(id string, timeout time.Duration) error {
	ch := w.WaitFor(id)
	select {
	case <-ch:
		return nil
	case <-time.After(timeout):
		return errors.New(codes.DeadlineExceeded)
	}
}

func (w *Waiter) WaitFor(id string) <-chan struct{} {
	w.mu.Lock()
	defer w.mu.Unlock()

	ch, exists := w.waitCh[id]
	if !exists {
		ch = make(chan struct{})
		w.waitCh[id] = ch
	}
	return ch
}

func (w *Waiter) Notify(id string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if ch, exists := w.waitCh[id]; exists {
		close(ch)
		delete(w.waitCh, id)
	}
}
