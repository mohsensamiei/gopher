package slices

func Remove[T comparable](values []T, drop ...T) []T {
	var result []T
	for _, item := range values {
		if Contains(item, drop...) {
			continue
		}
		result = append(result, item)
	}
	return result
}
