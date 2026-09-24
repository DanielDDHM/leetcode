package main

import (
	"testing"
)

func buildList(values [][]int) *Node {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*Node, len(values))
	for i := range nodes {
		nodes[i] = &Node{Val: values[i][0]}
	}

	for i := 0; i < len(values)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	for i, val := range values {
		if val[1] != -1 {
			nodes[i].Random = nodes[val[1]]
		}
	}

	return nodes[0]
}

func verifyCopy(t *testing.T, original, copy *Node) {
	if original == nil && copy == nil {
		return
	}
	if (original == nil) != (copy == nil) {
		t.Errorf("one is nil, the other is not")
		return
	}
	if original == copy {
		t.Errorf("copy points to same node as original")
		return
	}

	nodeMap := make(map[*Node]*Node)
	curr := original
	for curr != nil {
		nodeMap[curr] = curr
		curr = curr.Next
	}

	currOrig := original
	currCopy := copy
	for currOrig != nil {
		if currCopy == nil {
			t.Errorf("copy list shorter than original")
			return
		}
		if currOrig == currCopy {
			t.Errorf("copy node is same as original node")
			return
		}
		if currOrig.Val != currCopy.Val {
			t.Errorf("value mismatch: %d vs %d", currOrig.Val, currCopy.Val)
			return
		}

		if (currOrig.Next == nil) != (currCopy.Next == nil) {
			t.Errorf("next pointer mismatch at value %d", currOrig.Val)
			return
		}

		if (currOrig.Random == nil) != (currCopy.Random == nil) {
			t.Errorf("random pointer mismatch at value %d", currOrig.Val)
			return
		}

		if currOrig.Next != nil && currCopy.Next.Val != currOrig.Next.Val {
			t.Errorf("next pointer value mismatch at value %d", currOrig.Val)
			return
		}

		if currOrig.Random != nil && currCopy.Random.Val != currOrig.Random.Val {
			t.Errorf("random pointer value mismatch at value %d", currOrig.Val)
			return
		}

		if currCopy.Next != nil && nodeMap[currCopy.Next] != nil {
			t.Errorf("copy next pointer points to original node")
			return
		}

		if currCopy.Random != nil && nodeMap[currCopy.Random] != nil {
			t.Errorf("copy random pointer points to original node")
			return
		}

		currOrig = currOrig.Next
		currCopy = currCopy.Next
	}

	if currCopy != nil {
		t.Errorf("copy list longer than original")
		return
	}
}

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name   string
		values [][]int
	}{
		{
			name:   "example1",
			values: [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		{
			name:   "example2",
			values: [][]int{{1, 1}, {2, 1}},
		},
		{
			name:   "example3",
			values: [][]int{{3, -1}, {3, 0}, {3, -1}},
		},
		{
			name:   "empty list",
			values: [][]int{},
		},
		{
			name:   "single node with null random",
			values: [][]int{{5, -1}},
		},
		{
			name:   "single node with self random",
			values: [][]int{{5, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildList(tt.values)
			copy := copyRandomList(original)
			verifyCopy(t, original, copy)
		})
	}
}
