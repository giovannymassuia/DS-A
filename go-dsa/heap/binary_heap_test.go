package heap

import (
	"testing"
)

func TestMaxBinaryHeap(t *testing.T) {
	heap := NewBinaryHeap(10, MaxHeap)

	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(15)
	heap.Insert(40)
	heap.Insert(50)
	heap.Insert(100)
	heap.Insert(25)
	heap.Insert(45)

	// Array: [100, 45, 50, 40, 20, 15, 25, 10]
	/* Tree representation
		 100
	   /     \
	  45	  50
	 /  \    /  \
	40  20  15  25
	*/

	tests := []struct {
		name         string
		expectResult int
		expectArray  []int
	}{
		{"Peek", 100, []int{100, 45, 50, 40, 20, 15, 25, 10}},
		/* Tree after peek
			 100
		   /     \
		  45	  50
		 /  \    /  \
		40  20  15  25
		*/
		{"Pop", 100, []int{50, 45, 25, 40, 20, 15, 10}},
		/* Tree after pop
			 50
		   /     \
		  45      25
		 /  \    /  \
		40  20  15  10
		*/
		{"Pop", 50, []int{45, 40, 25, 10, 20, 15}},
		/* Tree after pop
			 45
		   /     \
		  40      25
		 /  \    /  \
		10  20  15
		*/
		{"Delete", 45, []int{45, 20, 25, 10, 15}},
		/* Tree after delete
			 45
		   /     \
		  20      25
		 /  \    /
		10  15
		*/
		{"Peek", 45, []int{45, 20, 25, 10, 15}},
		/* Tree after peek
			 45
		   /     \
		  20      25
		 /  \    /
		10  15
		*/
		{"Pop", 45, []int{25, 20, 15, 10}},
		/* Tree after pop
			 25
		   /     \
		  20      15
		 /
		10
		*/
		{"Pop", 25, []int{20, 10, 15}},
		/* Tree after pop
			 20
		   /     \
		  10      15
		*/
	}

	// should run tests in order they are defined
	runTests(t, tests, heap, 40)
}

func TestMinBinaryHeap(t *testing.T) {
	heap := NewBinaryHeap(10, MinHeap)

	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(15)
	heap.Insert(40)
	heap.Insert(50)
	heap.Insert(100)
	heap.Insert(25)
	heap.Insert(45)

	// Array: [10, 20, 15, 40, 50, 100, 25, 45]
	/* Tree representation
		 10
	   /     \
	  20 	  15
	 /  \    /  \
	40  50  100  25
	*/

	tests := []struct {
		name         string
		expectResult int
		expectArray  []int
	}{
		{"Peek", 10, []int{10, 20, 15, 40, 50, 100, 25, 45}},
		{"Pop", 10, []int{15, 20, 25, 40, 50, 100, 45}},
		{"Pop", 15, []int{20, 40, 25, 45, 50, 100}},
		{"Delete", 25, []int{25, 40, 100, 45, 50}},
		{"Peek", 25, []int{25, 40, 100, 45, 50}},
		{"Pop", 25, []int{40, 45, 100, 50}},
		{"Pop", 40, []int{45, 50, 100}},
	}

	// should run tests in order they are defined
	runTests(t, tests, heap, 20)
}

func TestHeapifyArray(t *testing.T) {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap := HeapifyArray(arr, MinHeap)

	expectedArr := []int{1, 2, 4, 3, 6, 5, 8, 10, 7, 9}
	/* Tree representation
				  1
			   /     \
			  2 	  4
			 /  \    /  \
			3    6  5    8
	       /  \
	      10   7
	*/

	if !compareArray(heap.GetData(), expectedArr) {
		t.Errorf("Expected %v but got %v", expectedArr, heap.GetData())
	}
}

func runTests(t *testing.T, tests []struct {
	name         string
	expectResult int
	expectArray  []int
}, heap *BinaryHeap, deleteValue int) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int
			if test.name == "Peek" {
				result = heap.Peek()
			} else if test.name == "Pop" {
				result = heap.Pop()
			} else if test.name == "Delete" {
				heap.Delete(deleteValue)
				result = heap.Peek()
			}

			if result != test.expectResult {
				t.Errorf("Expected %d but got %d", test.expectResult, result)
			}
			if !compareArray(heap.GetData(), test.expectArray) {
				t.Errorf("Expected %v but got %v", test.expectArray, heap.GetData())
			}
		})
	}
}

func compareArray(actual, expected []int) bool {
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			return false
		}
	}

	return true
}
