package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if clone, exists := visited[node]; exists {
		return clone
	}
	clone := &Node{Val: node.Val}
	visited[node] = clone
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(neighbor, visited))
	}
	return clone
}
