package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

type TestData struct {
	input    []int
	expected []int
}

func TestSorting(t *testing.T) {
	algorithms := []Sorting{
		&InsertionSort{},
		&HeapSort{},
		&ShellSort{},
	}

	tests := buildTestData()

	runtTests(t, algorithms, tests)
}

func runtTests(t *testing.T, algorithms []Sorting, tests []TestData) {
	for _, algo := range algorithms {
		algoName := reflect.TypeOf(algo).Elem().Name()
		t.Run(algoName, func(t *testing.T) {
			for idx, test := range tests {
				testName := fmt.Sprintf("Test %d", idx)
				t.Run(testName, func(t *testing.T) {
					arr := make([]int, len(test.input))
					copy(arr, test.input)

					algo.Sort(arr)

					if !compareArray(arr, test.expected) {
						t.Errorf("Expected %v but got %v", test.expected, arr)
					}
				})
			}
		})
	}
}

func buildTestData() []TestData {
	return []TestData{
		{[]int{10, 20, 15, 40, 50, 100, 25, 45}, []int{10, 15, 20, 25, 40, 45, 50, 100}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{10, 20, 15, 40, 50, 100, 25, 45, 5, 35}, []int{5, 10, 15, 20, 25, 35, 40, 45, 50, 100}},
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
