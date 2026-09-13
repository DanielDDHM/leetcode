package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	var dfs func(inStart, inEnd, postStart, postEnd int) *TreeNode
	dfs = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
		if inStart > inEnd || postStart > postEnd {
			return nil
		}

		rootVal := postorder[postEnd]
		rootInIdx := indexMap[rootVal]

		leftSize := rootInIdx - inStart
		node := &TreeNode{Val: rootVal}
		node.Left = dfs(inStart, rootInIdx-1, postStart, postStart+leftSize-1)
		node.Right = dfs(rootInIdx+1, inEnd, postStart+leftSize, postEnd-1)

		return node
	}

	return dfs(0, len(inorder)-1, 0, len(postorder)-1)
}
