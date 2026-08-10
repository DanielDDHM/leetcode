package main

import (
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
	}{
		{
			name:     "example1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			name:     "example2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
		{
			name:     "empty_array",
			nums:     []int{},
			val:      1,
			expected: 0,
		},
		{
			name:     "all_elements_equal_to_val",
			nums:     []int{1, 1, 1, 1},
			val:      1,
			expected: 0,
		},
		{
			name:     "no_elements_equal_to_val",
			nums:     []int{1, 2, 3, 4},
			val:      5,
			expected: 4,
		},
		{
			name:     "single_element_equal_to_val",
			nums:     []int{1},
			val:      1,
			expected: 0,
		},
		{
			name:     "single_element_not_equal_to_val",
			nums:     []int{1},
			val:      2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			k := removeElement(numsCopy, tt.val)

			if k != tt.expected {
				t.Errorf("expected k=%d, got k=%d", tt.expected, k)
			}

			sort.Ints(numsCopy[:k])
			for i := 0; i < k; i++ {
				if numsCopy[i] == tt.val {
					t.Errorf("found val=%d at index %d in first %d elements", tt.val, i, k)
				}
			}
		})
	}
}
