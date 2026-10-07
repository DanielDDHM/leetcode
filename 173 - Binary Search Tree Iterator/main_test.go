package main

import (
	"testing"
)

func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestBSTIterator(t *testing.T) {
	tests := []struct {
		name   string
		vals   []interface{}
		ops    []string
		expect []interface{}
	}{
		{
			name: "example1",
			vals: []interface{}{7, 3, 15, nil, nil, 9, 20},
			ops: []string{
				"next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext",
			},
			expect: []interface{}{3, 7, true, 9, true, 15, true, 20, false},
		},
		{
			name:   "single_node",
			vals:   []interface{}{1},
			ops:    []string{"hasNext", "next", "hasNext"},
			expect: []interface{}{true, 1, false},
		},
		{
			name:   "left_spine",
			vals:   []interface{}{3, 2, nil, 1},
			ops:    []string{"next", "next", "next", "hasNext"},
			expect: []interface{}{1, 2, 3, false},
		},
		{
			name:   "right_spine",
			vals:   []interface{}{1, nil, 2, nil, 3},
			ops:    []string{"next", "next", "next", "hasNext"},
			expect: []interface{}{1, 2, 3, false},
		},
		{
			name:   "balanced_tree",
			vals:   []interface{}{5, 3, 7, 2, 4, 6, 8},
			ops:    []string{"next", "next", "next", "next", "next", "next", "next"},
			expect: []interface{}{2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.vals)
			it := Constructor(root)

			for i, op := range tt.ops {
				if op == "next" {
					result := it.Next()
					expectedVal := tt.expect[i].(int)
					if result != expectedVal {
						t.Errorf("next() = %d, want %d", result, expectedVal)
					}
				} else if op == "hasNext" {
					result := it.HasNext()
					expectedVal := tt.expect[i].(bool)
					if result != expectedVal {
						t.Errorf("hasNext() = %v, want %v", result, expectedVal)
					}
				}
			}
		})
	}
}
