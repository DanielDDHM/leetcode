package main

import (
	"testing"
)

func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	idx := 1
	for len(queue) > 0 && idx < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if vals[idx] != nil {
			node.Left = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Left)
		}
		idx++
		if idx < len(vals) && vals[idx] != nil {
			node.Right = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Right)
		}
		idx++
	}
	return root
}

func linkedListToSlice(node *TreeNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		if node.Left != nil {
			return nil
		}
		node = node.Right
	}
	return result
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name   string
		input  []interface{}
		expect []int
	}{
		{
			name:   "example 1",
			input:  []interface{}{1, 2, 5, 3, 4, nil, 6},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "example 2",
			input:  []interface{}{},
			expect: []int{},
		},
		{
			name:   "example 3",
			input:  []interface{}{0},
			expect: []int{0},
		},
		{
			name:   "only left child",
			input:  []interface{}{1, 2},
			expect: []int{1, 2},
		},
		{
			name:   "only right child",
			input:  []interface{}{1, nil, 2},
			expect: []int{1, 2},
		},
		{
			name:   "left-only chain",
			input:  []interface{}{1, 2, nil, 3},
			expect: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			flatten(root)
			result := linkedListToSlice(root)
			if len(result) != len(tt.expect) {
				t.Errorf("length mismatch: got %d, want %d", len(result), len(tt.expect))
				return
			}
			for i := range result {
				if result[i] != tt.expect[i] {
					t.Errorf("value mismatch at index %d: got %d, want %d", i, result[i], tt.expect[i])
				}
			}
		})
	}
}
