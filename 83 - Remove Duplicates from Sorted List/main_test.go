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

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"example1", []int{1, 1, 2}, []int{1, 2}},
		{"example2", []int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all duplicates", []int{5, 5, 5, 5, 5}, []int{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := arrayToList(tt.input)
			result := deleteDuplicates(head)
			output := listToArray(result)

			if len(output) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(output))
				return
			}

			for i := 0; i < len(output); i++ {
				if output[i] != tt.expected[i] {
					t.Errorf("at index %d: expected %d, got %d", i, tt.expected[i], output[i])
				}
			}
		})
	}
}
