package helper

func Compare_ints(a, b []int) bool {
	if len(a) < 1 || len(a) != len(b)  {
		return false
	}

	for i := 0 ; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
