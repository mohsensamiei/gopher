package pointers

func ToValue[T any](s *T) T {
	return *s
}

func ToPointer[T any](s T) *T {
	return &s
}
