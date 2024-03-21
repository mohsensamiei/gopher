package pointers

func Value[T comparable](s *T) T {
	if s == nil {
		var zero T
		return zero
	}
	return *s
}

func Pointer[T comparable](s T) *T {
	var zero T
	if s == zero {
		return nil
	}
	return &s
}
