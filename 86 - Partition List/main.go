package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{}
	greater := &ListNode{}
	lessPtr := less
	greaterPtr := greater

	for current := head; current != nil; current = current.Next {
		if current.Val < x {
			lessPtr.Next = current
			lessPtr = lessPtr.Next
		} else {
			greaterPtr.Next = current
			greaterPtr = greaterPtr.Next
		}
	}

	greaterPtr.Next = nil
	lessPtr.Next = greater.Next
	return less.Next
}
