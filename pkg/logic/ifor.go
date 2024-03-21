package logic

func IFor[T any, R any](f func(v T) R, slice ...T) []R {
	var res []R
	for _, v := range slice {
		res = append(res, f(v))
	}
	return res
}
