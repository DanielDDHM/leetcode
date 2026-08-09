package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := minHeap{}
	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}

	if len(h) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy

	for len(h) > 0 {
		node := heap.Pop(&h).(*ListNode)
		current.Next = node
		current = current.Next

		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
	}

	return dummy.Next
}
