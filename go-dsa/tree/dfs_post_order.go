package tree

func DFSPostOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	DFSPostOrder(n.Left, f)
	DFSPostOrder(n.Right, f)
	f(n)
}
