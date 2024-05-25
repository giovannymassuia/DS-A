package tree

func BFS(n *Node, f func(*Node)) {
	// big o time complexity: O(n)
	// big o space complexity: O(n)
	if n == nil {
		return
	}

	q := []*Node{n}
	for len(q) > 0 {
		node := q[0] // Dequeue
		q = q[1:]    // Remove the head of the queue

		f(node)

		if node.Left != nil {
			q = append(q, node.Left) // Enqueue
		}

		if node.Right != nil {
			q = append(q, node.Right) // Enqueue
		}
	}
}
