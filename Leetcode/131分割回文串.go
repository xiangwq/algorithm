package Leetcode

import "fmt"

func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
			fmt.Println(i, j, f[i][j])
		}
	}

	var dfs func(int)
	spiltes := make([]string, 0)
	ans := make([][]string, 0)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), spiltes...))
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				spiltes = append(spiltes, s[i:j+1])
				dfs(j + 1)
				spiltes = spiltes[:len(spiltes)-1]
			}
		}
	}
	dfs(0)
	return ans
}
