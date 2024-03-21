package slices

func Common[T comparable](slice1 []T, slice2 []T, min int) bool {
	if min <= 0 {
		min = 1
	}
	common := 0
	for _, i := range slice1 {
		for _, j := range slice2 {
			if i == j {
				common += 1
			}
		}
	}
	return common >= min
}
