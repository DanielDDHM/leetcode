package main

import (
	"testing"
)

func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))
	for i := 0; i < len(adjList); i++ {
		nodes[i] = &Node{Val: i + 1}
	}

	for i := 0; i < len(adjList); i++ {
		for _, neighbor := range adjList[i] {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}

	return nodes[0]
}

func verifyClone(original *Node, cloned *Node, visited map[*Node]bool) bool {
	if visited == nil {
		visited = make(map[*Node]bool)
	}

	if original == nil && cloned == nil {
		return true
	}
	if original == nil || cloned == nil {
		return false
	}
	if original == cloned {
		return false
	}
	if visited[original] {
		return true
	}

	visited[original] = true

	if original.Val != cloned.Val {
		return false
	}
	if len(original.Neighbors) != len(cloned.Neighbors) {
		return false
	}

	for i := 0; i < len(original.Neighbors); i++ {
		if !verifyClone(original.Neighbors[i], cloned.Neighbors[i], visited) {
			return false
		}
	}

	return true
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "Example 1: graph with 4 nodes",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "Example 2: graph with single empty node",
			adjList: [][]int{{}},
		},
		{
			name:    "Edge case: nil graph",
			adjList: [][]int{},
		},
		{
			name:    "Edge case: graph with single node with self loop",
			adjList: [][]int{{1}},
		},
		{
			name:    "Edge case: graph with 2 nodes",
			adjList: [][]int{{2}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			cloned := cloneGraph(original)

			if !verifyClone(original, cloned, nil) {
				t.Errorf("cloneGraph failed for %s", tt.name)
			}
		})
	}
}
