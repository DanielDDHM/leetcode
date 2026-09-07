package main

import (
	"testing"
)

func listToArray(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func arrayToList(arr []int) *ListNode {
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

func equal(a, b []int) bool {
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

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		left     int
		right    int
		expected []int
	}{
		{
			name:     "example 1",
			input:    []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			name:     "example 2",
			input:    []int{5},
			left:     1,
			right:    1,
			expected: []int{5},
		},
		{
			name:     "reverse entire list",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    5,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "two elements reverse",
			input:    []int{1, 2},
			left:     1,
			right:    2,
			expected: []int{2, 1},
		},
		{
			name:     "single element",
			input:    []int{1},
			left:     1,
			right:    1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := arrayToList(tt.input)
			result := reverseBetween(head, tt.left, tt.right)
			output := listToArray(result)

			if !equal(output, tt.expected) {
				t.Errorf("got %v, want %v", output, tt.expected)
			}
		})
	}
}
