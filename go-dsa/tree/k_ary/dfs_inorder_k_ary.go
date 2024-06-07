package k_ary

/*
DFSInOrder traverses the tree in in-order fashion.
In this case we need to define a rule to adapt the traversal to the k-ary tree.

In this example, we will recursively traverse half of the children, then visit the node,
and then traverse the other half of the children.
*/
func DFSInOrder(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	// Traverse the first half of the children
	for i := 0; i < len(n.Children)/2; i++ {
		DFSInOrder(n.Children[i], f)
	}

	// Visit the node
	f(n)

	// Traverse the second half of the children
	for i := len(n.Children) / 2; i < len(n.Children); i++ {
		DFSInOrder(n.Children[i], f)
	}

}
