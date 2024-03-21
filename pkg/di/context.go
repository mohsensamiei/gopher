package di

import "context"

type ToContext[T any] func(context.Context, T) context.Context

type FromContext[T any] func(context.Context) T
