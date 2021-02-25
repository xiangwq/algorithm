package main // import "algorithm"

import "fmt"

func main() {
	a := []int{1, 5, 3, 2}
	InsertSort(a)
}

func InsertSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	for i := 1; i < len(a); i++ {
		for j := i - 1; j >= 0; j-- {
			if a[j] < a[j+1] {
				break
			}
			a[j], a[j+1] = a[j+1], a[j]
		}
	}
	return a
}
