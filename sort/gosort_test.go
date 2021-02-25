package sort

import (
	"algorithm/helper"
	"testing"
)

func BenchmarkGoSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := helper.GenerateIntArray(1000)
		GoSort(a)
	}
}
