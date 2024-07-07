package heap

// BinaryHeap
// Arr[(i-1)/2]	Returns the parent node
// Arr[(2*i)+1]	Returns the left child node
// Arr[(2*i)+2]	Returns the right child node
type BinaryHeap struct {
	data     []int
	capacity int
	size     int
	heapType HeapType
}

func NewBinaryHeap(capacity int, heapType HeapType) *BinaryHeap {
	return &BinaryHeap{
		data:     make([]int, capacity),
		capacity: capacity,
		size:     0,
		heapType: heapType,
	}
}

func HeapifyArray(arr []int, heapType HeapType) *BinaryHeap {
	heap := &BinaryHeap{
		data:     arr,
		capacity: len(arr),
		size:     len(arr),
		heapType: heapType,
	}

	// heapify the array
	n := heap.size/2 - 1 // n is the last non-leaf node

	for i := n; i >= 0; i-- {
		heap.heapify(i)
	}

	return heap
}

func (h *BinaryHeap) GetData() []int {
	return h.data[:h.size]
}

func (h *BinaryHeap) Peek() int {
	if h.size == 0 {
		return -1
	}
	return h.data[0]
}

func (h *BinaryHeap) Pop() int {
	if h.size == 0 {
		return -1
	}

	// swap the first value with the last value
	swap(h.data, 0, h.size-1)
	h.size--

	// fix the heap property if it is violated
	h.heapify(0)

	return h.data[h.size]
}

func (h *BinaryHeap) heapify(index int) {
	leftIdx := getLeftChildIndex(index)
	rightIdx := getRightChildIndex(index)
	// refIndex is the index of the smallest element
	// if the heap is a min heap, otherwise it is the largest element
	refIndex := index

	if leftIdx < h.size && h.compare(h.data[leftIdx], h.data[refIndex]) {
		refIndex = leftIdx
	}

	if rightIdx < h.size && h.compare(h.data[rightIdx], h.data[refIndex]) {
		refIndex = rightIdx
	}

	if refIndex != index {
		swap(h.data, index, refIndex)
		h.heapify(refIndex)
	}
}

func (h *BinaryHeap) Insert(value int) bool {
	if h.size == h.capacity {
		return false
	}

	// first, insert the new element at the end
	h.data[h.size] = value
	h.size++

	// fix the min heap property if it is violated
	i := h.size - 1
	for i != 0 && !h.compare(h.data[getParentIndex(i)], h.data[i]) {
		swap(h.data, i, getParentIndex(i))
		i = getParentIndex(i)
	}

	return true
}

func (h *BinaryHeap) Delete(value int) {
	for i := 0; i < h.size; i++ {
		if h.data[i] == value {
			swap(h.data, i, h.size-1)
			h.size--
			h.heapify(i)
			break
		}
	}
}

func (h *BinaryHeap) compare(a, b int) bool {
	if h.heapType == MinHeap {
		return a < b
	}

	if h.heapType == MaxHeap {
		return a > b
	}

	return false
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(index int) int {
	return (2 * index) + 1
}

func getRightChildIndex(index int) int {
	return (2 * index) + 2
}
