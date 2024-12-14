package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to serve as the start of the merged list
	dummy := &ListNode{}
	current := dummy

	// While both lists are not empty
	for list1 != nil && list2 != nil {
		// Compare the values of the current nodes in each list
		if list1.Val <= list2.Val {
			// Add the node from list1 to the merged list
			current.Next = list1
			list1 = list1.Next
		} else {
			// Add the node from list2 to the merged list
			current.Next = list2
			list2 = list2.Next
		}
		// Move to the next node in the merged list
		current = current.Next
	}

	// If one of the lists is not empty, add the remaining elements to the merged list
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Return the merged list, starting from the next node of the dummy
	return dummy.Next
}

func main() {
	// Create two sample linked lists: list1 = [1, 2, 4], list2 = [1, 3, 4]
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	// Merge the lists
	mergedList := mergeTwoLists(list1, list2)

	// Print the merged list
	for mergedList != nil {
		fmt.Print(mergedList.Val, " -> ")
		mergedList = mergedList.Next
	}
	fmt.Println("nil")
}
