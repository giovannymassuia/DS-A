package tree

func DFSPreOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	f(n)
	DFSPreOrder(n.Left, f)
	DFSPreOrder(n.Right, f)
}
