package slices

func Clone[T any](src []T) []T {
	dest := make([]T, len(src))
	copy(dest, src)
	return dest
}
