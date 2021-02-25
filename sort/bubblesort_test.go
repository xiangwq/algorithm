package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	b := make([]int, 5)
	a := helper.GenerateIntArray(5)
	copy(b, a)
	BubbleSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("冒泡排序失败")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(5)
		BubbleSort(a)
	}
}
