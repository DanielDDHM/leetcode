package main

import "testing"

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

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name   string
		lists  [][]int
		expect []int
	}{
		{
			name:   "Example 1",
			lists:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:   "Example 2 - empty lists",
			lists:  [][]int{},
			expect: []int{},
		},
		{
			name:   "Example 3 - single empty list",
			lists:  [][]int{{}},
			expect: []int{},
		},
		{
			name:   "Single non-empty list",
			lists:  [][]int{{1, 2, 3}},
			expect: []int{1, 2, 3},
		},
		{
			name:   "Multiple lists with duplicates",
			lists:  [][]int{{1, 1}, {1, 1}},
			expect: []int{1, 1, 1, 1},
		},
		{
			name:   "Lists with negative numbers",
			lists:  [][]int{{-10, -5, 0}, {-8, -3, 1}},
			expect: []int{-10, -8, -5, -3, 0, 1},
		},
		{
			name:   "All empty lists",
			lists:  [][]int{{}, {}, {}},
			expect: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tt.lists))
			for i, vals := range tt.lists {
				lists[i] = sliceToList(vals)
			}

			result := mergeKLists(lists)
			got := listToSlice(result)

			if len(got) != len(tt.expect) {
				t.Errorf("length mismatch: got %d, want %d", len(got), len(tt.expect))
				return
			}

			for i := range got {
				if got[i] != tt.expect[i] {
					t.Errorf("at index %d: got %d, want %d", i, got[i], tt.expect[i])
				}
			}
		})
	}
}
