package slices

func Contains[T comparable](wanted T, slice ...T) bool {
	for _, item := range slice {
		if item == wanted {
			return true
		}
	}
	return false
}
