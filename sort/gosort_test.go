package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func BenchmarkGoSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(100)
		GoSort(a)
	}
}

func TestGoSort(t *testing.T) {
	b := make([]int, 100)
	a := helper.GenerateIntArray(100)
	copy(b, a)
	GoSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("go排序失败")
	}
}
