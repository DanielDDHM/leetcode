package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := findFirst(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}

	right := findLast(nums, target)
	return []int{left, right}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			result = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			result = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
