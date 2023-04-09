package slices

import "fmt"

func ToInterface[T any](values []T) []any {
	var results []any
	for _, val := range values {
		results = append(results, val)
	}
	return results
}

func ToString[T any](values []T) []string {
	var results []string
	for _, val := range values {
		results = append(results, fmt.Sprint(val))
	}
	return results
}
