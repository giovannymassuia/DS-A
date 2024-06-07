package k_ary

func DFSPostOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	for _, child := range n.Children {
		DFSPostOrder(child, f)
	}

	f(n)
}
