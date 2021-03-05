package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	b := make([]int, 10000)
	a := helper.GenerateIntArray(10000)
	copy(b, a)
	MergeSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("归并排序错误")
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(10000)
		MergeSort(a)
	}
}
