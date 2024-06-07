package k_ary

import (
	"reflect"
	"testing"
)

func TestDFSPostOrder(t *testing.T) {
	/*
			Should print numbers in sequence: 1, 2, 3, 4, 5, 6, 7, 8, 9
				Tree:
						  9
		             /          \
				    4            8
				 /  |  \      /  |  \
				1   2   3    5   6   7
	*/

	root := NewNode(9)
	root.Children = []*Node{
		NewNode(4),
		NewNode(8),
	}
	root.Children[0].Children = []*Node{
		NewNode(1),
		NewNode(2),
		NewNode(3),
	}
	root.Children[1].Children = []*Node{
		NewNode(5),
		NewNode(6),
		NewNode(7),
	}

	var result []int
	DFSPostOrder(root, func(n *Node) {
		result = append(result, n.Value)
	})

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
