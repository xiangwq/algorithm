package sort

func CountingSort(a []int) {
	if len(a) <= 1 {
		return
	}
	var min, max int
	min = a[0]
	max = a[0]

	for _, v := range a {
		if min > v {
			min = v
		}

		if max < v {
			max = v
		}

	}

	arr := make([]int, max-min+1, max-min+1)
	for _, v := range a {
		arr[v-min]++
	}
	index := 0
	for k, v := range arr {
		if v == 0 {
			continue
		}
		for i := 0; i < v; i++ {
			a[index] = k + min
			index++
		}
	}

}
