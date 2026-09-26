package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverse(slow.Next)
	slow.Next = nil

	merge(head, second)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func merge(l1, l2 *ListNode) {
	for l2 != nil {
		next1 := l1.Next
		next2 := l2.Next

		l1.Next = l2
		l2.Next = next1

		l1 = next1
		l2 = next2
	}
}
