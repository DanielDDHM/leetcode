package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var current []int
	visited := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			result = append(result, append([]int{}, current...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			visited[i] = true
			current = append(current, nums[i])
			backtrack()
			current = current[:len(current)-1]
			visited[i] = false
		}
	}

	backtrack()
	return result
}
