package k_ary

import (
	"reflect"
	"testing"
)

func buildOrderTree() *Node {
	/*
			Should print numbers in sequence: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
			Tree:
		                  1
					 /          \
					2             3
				 /  |  \      /  |  |  \
				4   5   6    7   8  9   10
	*/
	root := NewNodeWithChildren(1, []*Node{
		NewNodeWithChildren(2, []*Node{
			NewNode(4),
			NewNode(5),
			NewNode(6),
		}),
		NewNodeWithChildren(3, []*Node{
			NewNode(7),
			NewNode(8),
			NewNode(9),
			NewNode(10),
		}),
	})

	return root
}

func TestBFS(t *testing.T) {
	var result []int
	root := buildOrderTree()

	BFS(root, func(n *Node) {
		result = append(result, n.Value)
	})

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
