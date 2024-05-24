package tree

import (
	"reflect"
	"testing"
)

func TestDFSInOrder(t *testing.T) {
	/*
		Should print numbers in sequence: 1, 2, 3, 4, 5, 6
			Tree:
				4
			   /  \
			  2    5
			 / \     \
			1   3     6
	*/

	root := NewNode(4)
	root.Left = NewNode(2)
	root.Right = NewNode(5)
	root.Left.Left = NewNode(1)
	root.Left.Right = NewNode(3)
	root.Right.Right = NewNode(6)

	var result []int
	DFSInOrder(root, func(n *Node) {
		result = append(result, n.Value)
	})

	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
