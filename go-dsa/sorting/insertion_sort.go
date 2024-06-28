package sorting

type InsertionSort struct{}

func (i *InsertionSort) Sort(arr []int) {
	// 1. start with the second element of the array
	// 2. compare the current element with the previous element
	// 3. if the current element is smaller than the previous element, swap them
	// 4. repeat step 2 and 3 until the current element is greater than the previous element
	// 5. move to the next element and repeat steps 2 to 4
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}
