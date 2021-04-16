package sort

import (
	"algorithm/helper"
	"sort"
	"testing"
)

func TestBucketSort(t *testing.T) {
	b := make([]int, 10000)
	a := helper.GenerateIntArray(10000)
	copy(b, a)
	BucketSort(a)
	sort.Ints(b)
	if !helper.Compare_ints(a, b) {
		t.Error("计数排序失败")
	}
}

func BenchmarkBucketSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(10000)
		BucketSort(a)
	}
}
