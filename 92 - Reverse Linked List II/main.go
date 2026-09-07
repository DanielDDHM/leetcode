package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	prevLeft := dummy

	for i := 1; i < left; i++ {
		prevLeft = prevLeft.Next
	}

	leftNode := prevLeft.Next
	rightNode := leftNode
	for i := left; i < right; i++ {
		rightNode = rightNode.Next
	}

	nextRight := rightNode.Next
	rightNode.Next = nil

	var prev *ListNode
	current := leftNode
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	prevLeft.Next = prev
	leftNode.Next = nextRight

	return dummy.Next
}
