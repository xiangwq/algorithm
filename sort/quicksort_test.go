package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	b := make([]int, 10000)
	a := helper.GenerateIntArray(10000)
	copy(b, a)
	QuickSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("快速排序错误")
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(10000)
		QuickSort(a)
	}
}
