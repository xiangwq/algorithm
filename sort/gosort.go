package sort

import "sort"

func GoSort(a []int) []int {
	sort.Ints(a)
	return a
}
