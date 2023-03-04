package slices

import (
	"fmt"
	"math/rand"
	"time"
)

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

func Contains[T comparable](wanted T, slice ...T) bool {
	for _, item := range slice {
		if item == wanted {
			return true
		}
	}
	return false
}

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

func Shuffle[T any](slice []T) []T {
	var (
		dest   []T
		length = len(slice)
		src    = make([]T, length)
		random = rand.New(rand.NewSource(time.Now().Unix()))
	)
	copy(src, slice)
	for i := 0; i < length; i++ {
		index := random.Intn(len(src))
		dest = append(dest, src[index])
		src = append(src[:index], src[index+1:]...)
	}
	return dest
}
