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

func Keys[K comparable, V any](dic map[K]V) (keys []K) {
	for k, _ := range dic {
		keys = append(keys, k)
	}
	return
}

func Values[K comparable, V any](dic map[K]V) (values []V) {
	for _, v := range dic {
		values = append(values, v)
	}
	return
}
