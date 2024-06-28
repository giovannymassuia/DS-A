package sorting

type ShellSort struct{}

func (i *ShellSort) Sort(arr []int) {
	// 1. start with a large gap, then reduce the gap until it becomes 1
	// 2. perform insertion sort on the elements with the gap

	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		// start from the gap element and move to the right
		for i := gap; i < len(arr); i++ {
			// store the current element
			temp := arr[i]
			j := i

			// run the insertion sort on the elements with the gap
			// shift the elements to the right until the correct position is found
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}

			// insert the temp element at the correct position
			arr[j] = temp
		}
	}
}
