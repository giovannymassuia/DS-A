package tree

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	/*
		Should print numbers in sequence: 1, 2, 3, 4, 5, 6
			Tree:
				1
			   /  \
			  2    3
			 / \     \
			4   5     6
	*/

	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Right = NewNode(6)

	var result []int
	BFS(root, func(n *Node) {
		result = append(result, n.Value)
	})

	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
