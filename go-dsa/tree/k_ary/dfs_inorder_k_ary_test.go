package k_ary

import (
	"reflect"
	"testing"
)

func buildInOrderTree() *Node {
	/*
			Should print numbers in sequence: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
			Tree:
		                  5
					 /          \
					2             8
				 /  |  \      /  |  |  \
				1   3   4    6   7  9   10
	*/
	root := NewNodeWithChildren(5, []*Node{
		NewNodeWithChildren(2, []*Node{
			NewNode(1),
			NewNode(3),
			NewNode(4),
		}),
		NewNodeWithChildren(8, []*Node{
			NewNode(6),
			NewNode(7),
			NewNode(9),
			NewNode(10),
		}),
	})

	return root
}

func TestDFSInOrder(t *testing.T) {
	var result []int
	root := buildInOrderTree()

	DFSInOrder(root, func(n *Node) {
		result = append(result, n.Value)
	})

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
