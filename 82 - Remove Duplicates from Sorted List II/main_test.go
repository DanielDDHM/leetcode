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

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example1",
			input:    []int{1, 2, 3, 3, 4, 4, 5},
			expected: []int{1, 2, 5},
		},
		{
			name:     "example2",
			input:    []int{1, 1, 1, 2, 3},
			expected: []int{2, 3},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{},
		},
		{
			name:     "empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "duplicates at start",
			input:    []int{1, 1, 2, 3},
			expected: []int{2, 3},
		},
		{
			name:     "duplicates at end",
			input:    []int{1, 2, 3, 3},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			result := deleteDuplicates(head)
			got := listToSlice(result)
			if len(got) != len(tt.expected) {
				t.Errorf("got length %d, expected %d", len(got), len(tt.expected))
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("at index %d: got %d, expected %d", i, got[i], tt.expected[i])
				}
			}
		})
	}
}
