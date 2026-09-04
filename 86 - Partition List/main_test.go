package main

import (
	"testing"
)

func sliceToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for current := head; current != nil; current = current.Next {
		result = append(result, current.Val)
	}
	return result
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		head     []int
		x        int
		expected []int
	}{
		{
			name:     "Example 1",
			head:     []int{1, 4, 3, 2, 5, 2},
			x:        3,
			expected: []int{1, 2, 2, 4, 3, 5},
		},
		{
			name:     "Example 2",
			head:     []int{2, 1},
			x:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Empty list",
			head:     []int{},
			x:        5,
			expected: []int{},
		},
		{
			name:     "Single node less than x",
			head:     []int{1},
			x:        2,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.head)
			result := partition(head, tt.x)
			resultSlice := listToSlice(result)
			if len(resultSlice) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(resultSlice))
			}
			for i := range resultSlice {
				if resultSlice[i] != tt.expected[i] {
					t.Errorf("Expected %v, got %v", tt.expected, resultSlice)
					return
				}
			}
		})
	}
}
