package sorting

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{10, 20, 15, 40, 50, 100, 25, 45}, []int{10, 15, 20, 25, 40, 45, 50, 100}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{10, 20, 15, 40, 50, 100, 25, 45, 5, 35}, []int{5, 10, 15, 20, 25, 35, 40, 45, 50, 100}},
	}

	runTests(t, tests, &InsertionSort{})
}
