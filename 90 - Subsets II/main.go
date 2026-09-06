package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	backtrack(nums, 0, []int{}, &result)
	return result
}

func backtrack(nums []int, start int, current []int, result *[][]int) {
	subset := make([]int, len(current))
	copy(subset, current)
	*result = append(*result, subset)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		current = append(current, nums[i])
		backtrack(nums, i+1, current, result)
		current = current[:len(current)-1]
	}
}
