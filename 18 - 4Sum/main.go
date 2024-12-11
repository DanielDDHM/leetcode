package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	n := len(nums)
	if n < 4 {
		return result
	}

	// Sort the array
	sort.Ints(nums)

	// Iterate through the array
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates for the first number
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue // Skip duplicates for the second number
			}
			// Two-pointer approach
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					// Skip duplicates for the third and fourth numbers
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}
