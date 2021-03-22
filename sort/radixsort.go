package sort

func RadixSort(a []int) {
	if len(a) <= 1 {
		return
	}

	// 获取数据最大值
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	// 获取最大值位数
	count := 0
	for max/10 > 0 {
		max = max / 10
		count++
	}

	for i := 0; i <= count; i++ {
		temp := make([][]int, 10)
		for ti := range temp {
			temp[ti] = make([]int, 0)
		}

		for _, v := range a {
			index := locationNum(v, i)
			temp[index] = append(temp[index], v)
		}

		aindex := 0

		for _, v := range temp {
			for _, num := range v {
				a[aindex] = num
				aindex++
			}
		}
	}
}

func locationNum(num, index int) int {
	for i := 0; i < index; i++ {
		num = num / 10
	}
	return num % 10
}
