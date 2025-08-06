package mapext

func Merge[K comparable, V any](src, ext map[K]V) map[K]V {
	newMap := map[K]V{}
	for k, v := range src {
		newMap[k] = v
	}
	for k, v := range ext {
		newMap[k] = v
	}
	return newMap
}

func Keys[K comparable, V any](dic map[K]V) []K {
	var (
		i    = 0
		keys = make([]K, len(dic))
	)
	for k := range dic {
		keys[i] = k
		i++
	}
	return keys
}

func Values[K comparable, V any](dic map[K]V) []V {
	var (
		i      = 0
		values = make([]V, len(dic))
	)
	for _, v := range dic {
		values[i] = v
		i++
	}
	return values
}
