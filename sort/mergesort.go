package sort

func MergeSort(a []int) {
	if len(a) <= 1 {
		return
	}
	mergeSort(a, 0, len(a))
}

func mergeSort(a []int, start, end int) {
	if end-start <= 1 {
		return
	}

	length := end - start
	leftEnd := start + length/2
	mergeSort(a, start, leftEnd)
	mergeSort(a, leftEnd, end)

	// 将排序好的放在临时空间
	temp := make([]int, 0)
	var i = start
	var j = leftEnd
	for {
		if i == leftEnd && j == end {
			break
		}

		if i < leftEnd && j < end && a[i] < a[j] {
			temp = append(temp, a[i])
			i++
			continue
		}

		if i < leftEnd && j < end && a[j] < a[i] {
			temp = append(temp, a[j])
			j++
			continue
		}

		if i < leftEnd {
			temp = append(temp, a[i])
			i++
			continue
		}

		if j < end {
			temp = append(temp, a[j])
			j++
			continue
		}
	}

	for k := range temp {
		a[start+k] = temp[k]
	}

}
