package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	b := make([]int, 10000)
	a := helper.GenerateIntArray(10000)
	copy(b, a)
	RadixSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("选择排序错误")
	}
}

func BenchmarkRadixSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(10000)
		RadixSort(a)
	}
}
