package main

import "testing"

func buildListWithCycle(values []int, cyclePos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	var cycleNode *ListNode

	if cyclePos == 0 {
		cycleNode = head
	}

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
		if i == cyclePos {
			cycleNode = current
		}
	}

	if cyclePos >= 0 && cyclePos < len(values) {
		current.Next = cycleNode
	}

	return head
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name        string
		values      []int
		cyclePos    int
		expectedVal int
		hasCycle    bool
	}{
		{
			name:        "Example 1: cycle at index 1",
			values:      []int{3, 2, 0, -4},
			cyclePos:    1,
			expectedVal: 2,
			hasCycle:    true,
		},
		{
			name:        "Example 2: cycle at index 0",
			values:      []int{1, 2},
			cyclePos:    0,
			expectedVal: 1,
			hasCycle:    true,
		},
		{
			name:        "Example 3: no cycle",
			values:      []int{1},
			cyclePos:    -1,
			expectedVal: 0,
			hasCycle:    false,
		},
		{
			name:        "Empty list",
			values:      []int{},
			cyclePos:    -1,
			expectedVal: 0,
			hasCycle:    false,
		},
		{
			name:        "Single node with self-cycle",
			values:      []int{1},
			cyclePos:    0,
			expectedVal: 1,
			hasCycle:    true,
		},
		{
			name:        "Cycle at end",
			values:      []int{1, 2, 3, 4, 5},
			cyclePos:    4,
			expectedVal: 5,
			hasCycle:    true,
		},
		{
			name:        "Two nodes: no cycle",
			values:      []int{1, 2},
			cyclePos:    -1,
			expectedVal: 0,
			hasCycle:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListWithCycle(tt.values, tt.cyclePos)
			result := detectCycle(head)

			if tt.hasCycle {
				if result == nil {
					t.Errorf("expected cycle node with value %d, got nil", tt.expectedVal)
				} else if result.Val != tt.expectedVal {
					t.Errorf("expected node value %d, got %d", tt.expectedVal, result.Val)
				}
			} else {
				if result != nil {
					t.Errorf("expected nil (no cycle), got node with value %d", result.Val)
				}
			}
		})
	}
}
