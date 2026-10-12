package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev2 := 0
	prev1 := nums[0]

	for i := 1; i < len(nums); i++ {
		curr := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
