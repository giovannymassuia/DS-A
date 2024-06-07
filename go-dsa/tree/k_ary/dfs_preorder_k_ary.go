package k_ary

func DFSPreOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	f(n)

	for _, child := range n.Children {
		DFSPreOrder(child, f)
	}

}
