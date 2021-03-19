package Leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	indexs := make([]int, 0)
	for i, v := range intervals {
		if v[0] > newInterval[1] || v[1] < newInterval[0] {
			indexs = append(indexs, i)
			continue
		}

		if v[0] < newInterval[0] {
			newInterval[0] = v[0]
		}
		if v[1] > newInterval[1] {
			newInterval[1] = v[1]
		}
	}

	if len(indexs) == 0 {
		res = append(res, newInterval)
		return res
	}

	isInsert := false
	for _, v := range indexs {
		if !isInsert && newInterval[1] < intervals[v][0] {
			res = append(res, newInterval)
			isInsert = true
		}
		res = append(res, intervals[v])
	}
	if !isInsert {
		res = append(res, newInterval)
	}
	return res
}
