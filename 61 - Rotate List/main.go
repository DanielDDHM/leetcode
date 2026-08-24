package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	tail.Next = head

	current := head
	for i := 0; i < length-k-1; i++ {
		current = current.Next
	}

	newHead := current.Next
	current.Next = nil

	return newHead
}
