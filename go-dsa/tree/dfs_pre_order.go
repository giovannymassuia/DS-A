package tree

func DFSPreOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	f(n)
	DFSPreOrder(n.Left, f)
	DFSPreOrder(n.Right, f)
}

func DFSPreOrderIterative(n *Node, f func(*Node)) {
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
