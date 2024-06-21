package heap

// insert -> O(logN)
// extract / pop -> O(1) - max for max heap, min for min heap
// delete -> O(logN)
// heapify -> O(log N)

type Heap interface {
	Insert(int)
	Pop() int
	Peek() int
	Delete(int)
}

// heap type enum
type HeapType int

const (
	MinHeap HeapType = iota
	MaxHeap
)


