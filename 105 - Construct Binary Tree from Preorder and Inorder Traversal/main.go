package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}
	return buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inorderMap)
}

func buildTreeHelper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, inorderMap map[int]int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	rootIndex := inorderMap[rootVal]
	leftSize := rootIndex - inStart

	root.Left = buildTreeHelper(preorder, preStart+1, preStart+leftSize, inorder, inStart, rootIndex-1, inorderMap)
	root.Right = buildTreeHelper(preorder, preStart+leftSize+1, preEnd, inorder, rootIndex+1, inEnd, inorderMap)

	return root
}
