package main

import (
	"testing"
)

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name     string
		head     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			head:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			name:     "Example 2",
			head:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 2, 1, 4, 5},
		},
		{
			name:     "k equals length",
			head:     []int{1, 2, 3, 4},
			k:        4,
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "k is 1",
			head:     []int{1, 2, 3, 4},
			k:        1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "single node",
			head:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "empty list",
			head:     []int{},
			k:        1,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.head)
			result := reverseKGroup(head, tt.k)
			got := listToSlice(result)

			if !slicesEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
