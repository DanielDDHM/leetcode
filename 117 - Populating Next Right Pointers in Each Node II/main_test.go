package main

import (
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want []*Node
	}{
		{
			name: "example1",
			root: &Node{
				Val: 1,
				Left: &Node{
					Val:   2,
					Left:  &Node{Val: 4},
					Right: &Node{Val: 5},
				},
				Right: &Node{
					Val:   3,
					Right: &Node{Val: 7},
				},
			},
			want: nil,
		},
		{
			name: "empty",
			root: nil,
			want: nil,
		},
		{
			name: "single_node",
			root: &Node{Val: 1},
			want: nil,
		},
		{
			name: "only_left_children",
			root: &Node{
				Val: 1,
				Left: &Node{
					Val:  2,
					Left: &Node{Val: 3},
				},
			},
			want: nil,
		},
		{
			name: "only_right_children",
			root: &Node{
				Val: 1,
				Right: &Node{
					Val:   2,
					Right: &Node{Val: 3},
				},
			},
			want: nil,
		},
		{
			name: "perfect_tree",
			root: &Node{
				Val:   1,
				Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
				Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.root
			result := connect(root)

			if root == nil && result != nil {
				t.Errorf("Expected nil, got %v", result)
			}

			if root != nil && result != root {
				t.Errorf("Expected same root, got different pointer")
			}

			if root != nil {
				queue := []*Node{root}
				for len(queue) > 0 {
					size := len(queue)
					var nextLevel []*Node

					for i := 0; i < size; i++ {
						node := queue[i]

						if i < size-1 {
							if node.Next != queue[i+1] {
								t.Errorf("node.Next should point to next node in level")
							}
						} else {
							if node.Next != nil {
								t.Errorf("last node in level should have nil Next")
							}
						}

						if node.Left != nil {
							nextLevel = append(nextLevel, node.Left)
						}
						if node.Right != nil {
							nextLevel = append(nextLevel, node.Right)
						}
					}

					queue = nextLevel
				}
			}
		})
	}
}
