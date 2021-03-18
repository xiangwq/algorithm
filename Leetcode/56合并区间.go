package Leetcode

import "sort"

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-intervals

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}

	starts := make([]int, 0)
	ends := make([]int, 0)
	for _, v := range intervals {
		starts = append(starts, v[0])
		ends = append(ends, v[1])
	}
	sort.Ints(starts)
	sort.Ints(ends)

	f := make([][]int, 0)

	start := 1
	lastStart := 0
	for {
		if lastStart == n-1 || start >= n {
			temp := make([]int, 2)
			temp[0] = starts[lastStart]
			temp[1] = ends[start-1]
			f = append(f, temp)
			break
		}
		if starts[start] > ends[start-1] {
			temp := make([]int, 2)
			temp[0] = starts[lastStart]
			temp[1] = ends[start-1]
			f = append(f, temp)
			lastStart = start
		}
		start++
	}
	return f
}
