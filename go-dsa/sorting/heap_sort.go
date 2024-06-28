package sorting

type HeapSort struct{}

func (h *HeapSort) Sort(arr []int) {
	h.heapifyArray(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		h.heapify(arr, 0, i)
	}
}

// heapifyArray the array as a max heap
func (h *HeapSort) heapifyArray(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		h.heapify(arr, i, len(arr))
	}
}

// heapify
func (h *HeapSort) heapify(arr []int, i int, n int) {
	ref := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[ref] {
		ref = left
	}

	if right < n && arr[right] > arr[ref] {
		ref = right
	}

	if ref != i {
		arr[i], arr[ref] = arr[ref], arr[i]
		h.heapify(arr, ref, n)
	}
}
