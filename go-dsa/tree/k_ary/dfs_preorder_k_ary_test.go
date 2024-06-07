package k_ary

import (
	"reflect"
	"testing"
)

func buildPreOrderTree() *Node {
	/*
		Should print numbers in sequence: 1, 2, 3, 4, 5, 6, 7, 8, 9
		Tree:
					  1
				 /          \
				2            6
			 /  |  \      /  |  \
			3   4   5    7   8   9
	*/
	root := NewNodeWithChildren(1, []*Node{
		NewNodeWithChildren(2, []*Node{
			NewNode(3),
			NewNode(4),
			NewNode(5),
		}),
		NewNodeWithChildren(6, []*Node{
			NewNode(7),
			NewNode(8),
			NewNode(9),
		}),
	})

	return root
}

func TestDFSPreOrder(t *testing.T) {
	var result []int
	root := buildPreOrderTree()

	DFSPreOrder(root, func(n *Node) {
		result = append(result, n.Value)
	})

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
