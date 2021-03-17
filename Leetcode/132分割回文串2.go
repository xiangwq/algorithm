package Leetcode

import "math"

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。

func minCut(s string) int {
	n := len(s)
	if len(s) <= 1 {
		return 0
	}
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
		}
	}

	g := make([]int, n)
	for i := 0; i < n; i++ {
		if f[0][i] {
			continue
		}
		g[i] = math.MaxInt64
		for j := 1; j <= i; j++ {
			if f[j][i] && g[j-1]+1 < g[i] {
				g[i] = g[j-1] + 1
			}
		}
	}
	return g[n-1]
}
