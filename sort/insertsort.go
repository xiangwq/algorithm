package sort

func InsertSort(a []int) {
	if len(a) <= 1 {
		return
	}
	for i := 1; i < len(a); i++ {
		for j := i - 1; j >= 0; j-- {
			if a[j] < a[j+1] {
				break
			}
			a[j], a[j+1] = a[j+1], a[j]
		}
	}
}
