package main

import (
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name     string
		root     *Node
		validate func(*Node) bool
	}{
		{
			name: "perfect binary tree",
			root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 3,
					Left: &Node{
						Val: 6,
					},
					Right: &Node{
						Val: 7,
					},
				},
			},
			validate: func(root *Node) bool {
				if root == nil {
					return true
				}

				level := []*Node{root}

				for len(level) > 0 {
					nextLevel := []*Node{}

					for i := 0; i < len(level); i++ {
						if i < len(level)-1 {
							if level[i].Next != level[i+1] {
								return false
							}
						} else {
							if level[i].Next != nil {
								return false
							}
						}

						if level[i].Left != nil {
							nextLevel = append(nextLevel, level[i].Left)
						}
						if level[i].Right != nil {
							nextLevel = append(nextLevel, level[i].Right)
						}
					}

					level = nextLevel
				}

				return true
			},
		},
		{
			name:     "empty tree",
			root:     nil,
			validate: func(root *Node) bool { return root == nil },
		},
		{
			name: "single node",
			root: &Node{Val: 1},
			validate: func(root *Node) bool {
				return root != nil && root.Next == nil
			},
		},
		{
			name: "two levels",
			root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
				},
				Right: &Node{
					Val: 3,
				},
			},
			validate: func(root *Node) bool {
				if root.Left == nil || root.Right == nil {
					return false
				}
				if root.Left.Next != root.Right {
					return false
				}
				if root.Right.Next != nil {
					return false
				}
				if root.Left.Left != nil || root.Left.Right != nil {
					return false
				}
				if root.Right.Left != nil || root.Right.Right != nil {
					return false
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := connect(tt.root)
			if !tt.validate(result) {
				t.Errorf("connect() failed validation")
			}
		})
	}
}
