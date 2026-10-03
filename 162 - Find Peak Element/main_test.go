package main

import "testing"

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
		isPeak   func([]int, int) bool
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 1},
			expected: 2,
			isPeak: func(nums []int, idx int) bool {
				if idx == 0 {
					return idx+1 < len(nums) && nums[idx] > nums[idx+1]
				}
				if idx == len(nums)-1 {
					return nums[idx] > nums[idx-1]
				}
				return nums[idx] > nums[idx-1] && nums[idx] > nums[idx+1]
			},
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			isPeak: func(nums []int, idx int) bool {
				if idx == 0 {
					return idx+1 < len(nums) && nums[idx] > nums[idx+1]
				}
				if idx == len(nums)-1 {
					return nums[idx] > nums[idx-1]
				}
				return nums[idx] > nums[idx-1] && nums[idx] > nums[idx+1]
			},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 0,
			isPeak: func(nums []int, idx int) bool {
				return true
			},
		},
		{
			name:     "Ascending array",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 4,
			isPeak: func(nums []int, idx int) bool {
				if idx == 0 {
					return idx+1 < len(nums) && nums[idx] > nums[idx+1]
				}
				if idx == len(nums)-1 {
					return nums[idx] > nums[idx-1]
				}
				return nums[idx] > nums[idx-1] && nums[idx] > nums[idx+1]
			},
		},
		{
			name:     "Descending array",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 0,
			isPeak: func(nums []int, idx int) bool {
				if idx == 0 {
					return idx+1 < len(nums) && nums[idx] > nums[idx+1]
				}
				if idx == len(nums)-1 {
					return nums[idx] > nums[idx-1]
				}
				return nums[idx] > nums[idx-1] && nums[idx] > nums[idx+1]
			},
		},
		{
			name:     "Two elements ascending",
			nums:     []int{1, 2},
			expected: 1,
			isPeak: func(nums []int, idx int) bool {
				if idx == 0 {
					return idx+1 < len(nums) && nums[idx] > nums[idx+1]
				}
				if idx == len(nums)-1 {
					return nums[idx] > nums[idx-1]
				}
				return nums[idx] > nums[idx-1] && nums[idx] > nums[idx+1]
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElement(tt.nums)
			if tt.expected > 0 && result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
				return
			}
			if !tt.isPeak(tt.nums, result) {
				t.Errorf("result %d is not a peak in %v", result, tt.nums)
			}
		})
	}
}
