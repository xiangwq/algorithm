package sort

func BucketSort(a []int) {
	if len(a) <= 1 {
		return
	}

	bucketSize := 10

	min := a[0]
	max := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// 获取区间跨度
	span := (max - min) / (bucketSize - 1)
	if ((max-min)%(bucketSize-1) != 0) || span == 0 {
		span += 1
	}
	buckets := make([][]int, bucketSize)

	for k := range buckets {
		buckets[k] = make([]int, 0)
	}

	// 将数据放入对应的桶中
	for _, v := range a {
		index := (v - min) / span
		buckets[index] = append(buckets[index], v)
	}
	// 对桶中的数据进行排序
	for _, v := range buckets {
		quickSort(v, 0, len(v))
	}

	index := 0
	for _, bucket := range buckets {
		for _, v := range bucket {
			a[index] = v
			index++
		}
	}
}

func quickSort(a []int, start, end int) {
	if end-start <= 1 {
		return
	}

	index := start
	endindex := end - 1
	for {
		if index == endindex {
			break
		}
		if a[index] >= a[index+1] {
			a[index], a[index+1] = a[index+1], a[index]
			index++
			continue
		}

		if a[index] <= a[index+1] {
			a[index+1], a[endindex] = a[endindex], a[index+1]
			endindex--
			continue
		}
	}
	quickSort(a, start, index)
	quickSort(a, index+1, end)

}
