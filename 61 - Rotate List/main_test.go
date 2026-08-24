package main

import (
	"testing"
)

func makeList(arr []int) *ListNode {
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
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{
			name: "example1",
			head: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "example2",
			head: []int{0, 1, 2},
			k:    4,
			want: []int{2, 0, 1},
		},
		{
			name: "empty list",
			head: []int{},
			k:    1,
			want: []int{},
		},
		{
			name: "single node",
			head: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "k equals zero",
			head: []int{1, 2, 3},
			k:    0,
			want: []int{1, 2, 3},
		},
		{
			name: "k equals length",
			head: []int{1, 2, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "two nodes rotate by one",
			head: []int{1, 2},
			k:    1,
			want: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := makeList(tt.head)
			result := rotateRight(head, tt.k)
			got := listToSlice(result)

			if len(got) != len(tt.want) {
				t.Errorf("length mismatch: got %v, want %v", got, tt.want)
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
