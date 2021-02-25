package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	b := make([]int, 100)
	a := helper.GenerateIntArray(100)
	copy(b, a)
	HeapSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("插入排序失败")
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(100)
		HeapSort(a)
	}
}
