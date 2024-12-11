package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node to simplify edge cases
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	// Move the first pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move both pointers until first reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node from the end
	second.Next = second.Next.Next

	// Return the updated list
	return dummy.Next
}
