package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	b := make([]int, 100)
	a := helper.GenerateIntArray(100)
	copy(b, a)
	SelectionSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("选择排序错误")
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(100)
		SelectionSort(a)
	}
}
