package slices

func Prepend[T any](s []T, v T) []T {
	return append([]T{v}, s...)
}
