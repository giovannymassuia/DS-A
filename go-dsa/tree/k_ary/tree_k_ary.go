package k_ary

type Node struct {
	Value    int
	Children []*Node
}

func NewNode(value int) *Node {
	return &Node{
		Value:    value,
		Children: []*Node{},
	}
}

func NewNodeWithChildren(value int, children []*Node) *Node {
	return &Node{
		Value:    value,
		Children: children,
	}
}
