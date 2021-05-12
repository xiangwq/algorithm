package Leetcode

func twoSum(nums []int, target int) []int {
	history := make(map[int]int, 0)
	for k, v := range nums {
		if index, ok := history[target-v]; ok {
			return []int{index, k}
		}
		history[v] = k
	}
	return []int{0, 0}
}
