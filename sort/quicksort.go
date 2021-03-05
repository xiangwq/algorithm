package sort

func QuickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	qucikSort(a, 0, len(a))

}

func qucikSort(a []int, start, end int) {
	if end-start <= 1 {
		return
	}

	index := start
	j := end - 1
	for {
		if index == j {
			break
		}
		if a[index] > a[index+1] {
			a[index], a[index+1] = a[index+1], a[index]
			index++
			continue
		}

		if a[index] <= a[index+1] {
			a[index+1], a[j] = a[j], a[index+1]
			j--
			continue
		}
	}
	qucikSort(a, start, index)
	qucikSort(a, index+1, end)
}
