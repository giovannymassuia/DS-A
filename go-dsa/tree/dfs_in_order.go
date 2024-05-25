package tree

func DFSInOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	DFSInOrder(n.Left, f)
	f(n)
	DFSInOrder(n.Right, f)
}

func DFSInOrderIterative(n *Node, f func(*Node)) {
	// big O space: O(h) where h is the height of the tree
	// big O time: O(n) where n is the number of nodes in the tree
	var stack []*Node
	current := n

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		current = current.Right
	}
}
