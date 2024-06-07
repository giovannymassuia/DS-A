package k_ary

func BFS(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	q := []*Node{n}
	for len(q) > 0 {
		node := q[0] // Dequeue
		q = q[1:]    // Remove the head of the queue

		f(node)

		for _, child := range node.Children {
			q = append(q, child) // Enqueue
		}
	}
}
