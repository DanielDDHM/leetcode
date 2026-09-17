package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			current := queue[i]

			if i < size-1 {
				current.Next = queue[i+1]
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		queue = queue[size:]
	}

	return root
}
