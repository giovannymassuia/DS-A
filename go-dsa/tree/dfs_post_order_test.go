package tree

import (
	"reflect"
	"testing"
)

func TestDFSPostOrder(t *testing.T) {
	/*
		Should print numbers in sequence: 1, 2, 3, 4, 5, 6
			Tree:
				6
			   /  \
			  3    5
			 / \     \
			1   2     4
	*/

	root := NewNode(6)
	root.Left = NewNode(3)
	root.Right = NewNode(5)
	root.Left.Left = NewNode(1)
	root.Left.Right = NewNode(2)
	root.Right.Right = NewNode(4)

	var result []int
	DFSPostOrder(root, func(n *Node) {
		result = append(result, n.Value)
	})

	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
