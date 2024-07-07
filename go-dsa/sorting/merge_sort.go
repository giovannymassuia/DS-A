package sorting

type MergeSort struct{}

func (i *MergeSort) Sort(arr []int) {
	// Divide: Divide the list or array recursively into two halves until it can no more be divided.
	// Conquer: Each subarray is sorted individually using the merge sort algorithm.
	// Merge: The sorted subarrays are merged back together in sorted order.
	//		  The process continues until all elements from both subarrays have been merged.
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	// Create temporary arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temporary arrays
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = arr[m+1+i]
	}

	// Merge the temporary arrays back into arr[l..r]
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of L[] if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of R[] if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
