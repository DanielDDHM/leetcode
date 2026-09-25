package main

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{
			name: "cycle at second node",
			head: func() *ListNode {
				node1 := &ListNode{Val: 3}
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 0}
				node4 := &ListNode{Val: -4}
				node1.Next = node2
				node2.Next = node3
				node3.Next = node4
				node4.Next = node2
				return node1
			}(),
			want: true,
		},
		{
			name: "single node with cycle",
			head: func() *ListNode {
				node := &ListNode{Val: 1}
				node.Next = node
				return node
			}(),
			want: true,
		},
		{
			name: "no cycle single node",
			head: &ListNode{Val: 1},
			want: false,
		},
		{
			name: "two nodes with cycle",
			head: func() *ListNode {
				node1 := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node1.Next = node2
				node2.Next = node1
				return node1
			}(),
			want: true,
		},
		{
			name: "two nodes no cycle",
			head: func() *ListNode {
				node1 := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node1.Next = node2
				return node1
			}(),
			want: false,
		},
		{
			name: "cycle at end",
			head: func() *ListNode {
				node1 := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				node1.Next = node2
				node2.Next = node3
				node3.Next = node3
				return node1
			}(),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
