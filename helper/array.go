package helper

import "math/rand"

func GenerateIntArray(length int) []int {
	a := make([]int, 0)
	for i := 0; i < length; i++ {
		a = append(a, rand.Intn(1<<5))
	}
	return a
}
