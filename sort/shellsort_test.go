package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func TestShellSort(t *testing.T) {
	b := make([]int, 100)
	a := helper.GenerateIntArray(100)
	copy(b, a)
	SelectionSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("希尔排序错误")
	}
}

func BenchmarkShellSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(100)
		ShellSort(a)
	}
}
