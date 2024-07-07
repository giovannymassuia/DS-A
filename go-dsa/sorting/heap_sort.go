package sorting

type HeapSort struct {
	reverse bool
}

// Sort the array using heap sort, reverse is optional
func (h *HeapSort) Sort(arr []int) {

	// build max heap for ascending order
	// build min heap for descending order
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyIterative(arr, i, len(arr))
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		h.heapifyIterative(arr, 0, i)
	}
}

// heapify
func (h *HeapSort) heapify(arr []int, i int, n int) {
	ref := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.compare(arr[left], arr[ref]) {
		ref = left
	}

	if right < n && h.compare(arr[right], arr[ref]) {
		ref = right
	}

	if ref != i {
		arr[i], arr[ref] = arr[ref], arr[i]
		h.heapify(arr, ref, n)
	}
}

// heapify iteratively
func (h *HeapSort) heapifyIterative(arr []int, i int, n int) {
	for {
		ref := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && h.compare(arr[left], arr[ref]) {
			ref = left
		}

		if right < n && h.compare(arr[right], arr[ref]) {
			ref = right
		}

		if ref == i {
			break
		}

		arr[i], arr[ref] = arr[ref], arr[i]
		i = ref
	}
}

func (h *HeapSort) compare(a, b int) bool {
	if h.reverse {
		return a < b
	}
	return a > b
}
