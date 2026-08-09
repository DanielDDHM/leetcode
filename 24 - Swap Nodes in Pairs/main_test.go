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

func sliceToList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example 1",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
		},
		{
			name:     "example 2",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "example 3",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "example 4",
			input:    []int{1, 2, 3},
			expected: []int{2, 1, 3},
		},
		{
			name:     "two nodes",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "five nodes",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 1, 4, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			result := swapPairs(head)
			got := listToSlice(result)

			if len(got) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, want %d", len(got), len(tt.expected))
				return
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("got %v, want %v", got, tt.expected)
					return
				}
			}
		})
	}
}
