package mediator

import (
	"sync"
)

func New() *Mediator {
	return &Mediator{}
}

type Mediator struct {
	requests      sync.Map
	notifications sync.Map
	behaviors     []any
	behaviorMutex sync.RWMutex
}
