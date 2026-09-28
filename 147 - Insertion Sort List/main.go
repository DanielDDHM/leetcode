package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	curr := head.Next
	prev := head

	for curr != nil {
		if prev.Val <= curr.Val {
			prev = curr
			curr = curr.Next
			continue
		}

		p := dummy
		for p.Next.Val < curr.Val {
			p = p.Next
		}

		prev.Next = curr.Next
		curr.Next = p.Next
		p.Next = curr
		curr = prev.Next
	}

	return dummy.Next
}
