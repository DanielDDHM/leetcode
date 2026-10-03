package main

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() (*ListNode, *ListNode, *ListNode)
		wantNode bool
	}{
		{
			name: "Example 1: intersection at node with value 8",
			setup: func() (*ListNode, *ListNode, *ListNode) {
				intersection := &ListNode{Val: 8}
				intersection.Next = &ListNode{Val: 4}
				intersection.Next.Next = &ListNode{Val: 5}

				headA := &ListNode{Val: 4}
				headA.Next = &ListNode{Val: 1}
				headA.Next.Next = intersection

				headB := &ListNode{Val: 5}
				headB.Next = &ListNode{Val: 6}
				headB.Next.Next = &ListNode{Val: 1}
				headB.Next.Next.Next = intersection

				return headA, headB, intersection
			},
			wantNode: true,
		},
		{
			name: "Example 2: intersection at node with value 2",
			setup: func() (*ListNode, *ListNode, *ListNode) {
				intersection := &ListNode{Val: 2}
				intersection.Next = &ListNode{Val: 4}

				headA := &ListNode{Val: 1}
				headA.Next = &ListNode{Val: 9}
				headA.Next.Next = &ListNode{Val: 1}
				headA.Next.Next.Next = intersection

				headB := &ListNode{Val: 3}
				headB.Next = intersection

				return headA, headB, intersection
			},
			wantNode: true,
		},
		{
			name: "Example 3: no intersection",
			setup: func() (*ListNode, *ListNode, *ListNode) {
				headA := &ListNode{Val: 2}
				headA.Next = &ListNode{Val: 6}
				headA.Next.Next = &ListNode{Val: 4}

				headB := &ListNode{Val: 1}
				headB.Next = &ListNode{Val: 5}

				return headA, headB, nil
			},
			wantNode: false,
		},
		{
			name: "Both lists are nil",
			setup: func() (*ListNode, *ListNode, *ListNode) {
				return nil, nil, nil
			},
			wantNode: false,
		},
		{
			name: "First list is nil",
			setup: func() (*ListNode, *ListNode, *ListNode) {
				headB := &ListNode{Val: 1}
				return nil, headB, nil
			},
			wantNode: false,
		},
		{
			name: "Second list is nil",
			setup: func() (*ListNode, *ListNode, *ListNode) {
				headA := &ListNode{Val: 1}
				return headA, nil, nil
			},
			wantNode: false,
		},
		{
			name: "Single node intersection",
			setup: func() (*ListNode, *ListNode, *ListNode) {
				intersection := &ListNode{Val: 1}

				headA := intersection
				headB := intersection

				return headA, headB, intersection
			},
			wantNode: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, expected := tt.setup()
			result := getIntersectionNode(headA, headB)
			if tt.wantNode && result != expected {
				t.Errorf("expected intersection node, got %v", result)
			}
			if !tt.wantNode && result != nil {
				t.Errorf("expected nil, got %v", result)
			}
		})
	}
}
