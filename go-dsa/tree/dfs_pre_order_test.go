package tree

import (
	"reflect"
	"testing"
)

func buildPreOrderTree() *Node {
	/*
		Tree:
			1
		   /  \
		  2    5
		 / \     \
		3   4     6
	*/
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(5)
	root.Left.Left = NewNode(3)
	root.Left.Right = NewNode(4)
	root.Right.Right = NewNode(6)

	return root
}

func TestDFSPreOrder(t *testing.T) {
	var result []int
	root := buildPreOrderTree()

	DFSPreOrder(root, func(n *Node) {
		result = append(result, n.Value)
	})

	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestDFSPreOrderIterative(t *testing.T) {
	var result []int
	root := buildPreOrderTree()

	expected := []int{1, 2, 3, 4, 5, 6}
	DFSPreOrderIterative(root, func(n *Node) {
		result = append(result, n.Value)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
