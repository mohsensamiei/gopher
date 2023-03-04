package pointers

func ToValue[T comparable](s *T) T {
	if s == nil {
		var zero T
		return zero
	}
	return *s
}

func ToPointer[T comparable](s T) *T {
	var zero T
	if s == zero {
		return nil
	}
	return &s
}
