package sort

/**
希尔排序
*/
func ShellSort(a []int) {
	if len(a) < 1 {
		return
	}

	for gap := len(a) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(a); i++ {
			var j = i
			for {
				if j-gap < 0 || a[j] > a[j-gap] {
					break
				}
				a[j], a[j-gap] = a[j-gap], a[j]
				j = j - gap
			}
		}
	}
}
