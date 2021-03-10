package Leetcode

func partition(str string) [][]string {
	l := len(str)
	final := make([][]string, 0)
	for i := 1; i <= l; i++ {
		if !isPalindrome(str[0:i]) {
			continue
		}
		after := getAfterStr(str[i:])
		if after == nil {
			final = append(final, []string{str[0:i]})
			continue
		}

		for _, v := range after {
			temp := make([]string, 1)
			temp[0] = str[0:i]
			temp = append(temp, v...)
			final = append(final, temp)
		}
	}
	return final
}

func isPalindrome(str string) bool {
	if len(str) < 1 {
		return false
	}
	l := len(str) / 2
	maxIndex := len(str) - 1
	for i := 0; i <= l; i++ {
		if str[i] != str[maxIndex-i] {
			return false
		}
	}
	return true
}

func getAfterStr(str string) [][]string {
	if len(str) < 1 {
		return nil
	}
	final := make([][]string, 0)
	for i := 1; i <= len(str); i++ {
		if !isPalindrome(str[0:i]) {
			continue
		}
		after := getAfterStr(str[i:])
		if after == nil {
			final = append(final, []string{str[0:i]})
			continue
		}

		for _, v := range after {
			temp := make([]string, 0)
			temp = append(temp, str[0:i])
			temp = append(temp, v...)
			final = append(final, temp)
		}
	}
	return final
}
