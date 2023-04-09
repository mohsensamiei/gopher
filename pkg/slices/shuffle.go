package slices

import (
	"math/rand"
	"time"
)

func Shuffle[T any](slice []T) []T {
	var (
		dest   []T
		src    = Clone(slice)
		random = rand.New(rand.NewSource(time.Now().Unix()))
	)
	for i := 0; i < len(slice); i++ {
		index := random.Intn(len(src))
		dest = append(dest, src[index])
		src = append(src[:index], src[index+1:]...)
	}
	return dest
}
