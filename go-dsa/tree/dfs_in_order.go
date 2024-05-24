package tree

func DFSInOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	DFSInOrder(n.Left, f)
	f(n)
	DFSInOrder(n.Right, f)
}
