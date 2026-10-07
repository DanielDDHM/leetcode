package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var stack []*TreeNode
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	return BSTIterator{
		stack: stack,
	}
}

func (this *BSTIterator) Next() int {
	cur := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	for node := cur.Right; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	return cur.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
