package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd := nums[0]
	minProd := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		newMax := max(num, max(maxProd*num, minProd*num))
		newMin := min(num, min(maxProd*num, minProd*num))

		maxProd = newMax
		minProd = newMin
		result = max(result, maxProd)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
