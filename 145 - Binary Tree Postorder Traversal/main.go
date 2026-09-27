package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack1 := []*TreeNode{root}
	var stack2 []*TreeNode

	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node)

		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	for i := len(stack2) - 1; i >= 0; i-- {
		result = append(result, stack2[i].Val)
	}

	return result
}
