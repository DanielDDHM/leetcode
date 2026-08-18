package main

func permute(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, &result, used)
	return result
}

func backtrack(nums []int, current []int, result *[][]int, used []bool) {
	if len(current) == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			current = append(current, nums[i])
			backtrack(nums, current, result, used)
			current = current[:len(current)-1]
			used[i] = false
		}
	}
}
