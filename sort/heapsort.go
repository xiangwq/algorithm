package sort

func HeapSort(a []int) {
	if len(a) <= 1 {
		return
	}
	buildMaxHeap(a)
	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		adjustHeap(a[0:i], 0)
	}
}

// 构建大顶堆
func buildMaxHeap(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		adjustHeap(a, i)
	}
}

func adjustHeap(a []int, pos int) {
	large := pos
	// 需要排序的长度
	length := len(a)
	left := 2*pos + 1  //当前节点对应的左孩子的节点
	right := 2*pos + 2 // 当前孩子对应的右节点的孩子
	if left < length && a[pos] < a[left] {
		large = left
	}

	if right < length && a[large] < a[right] {
		large = right
	}
	if large != pos {
		a[large], a[pos] = a[pos], a[large]
		adjustHeap(a, large)
	}
}
