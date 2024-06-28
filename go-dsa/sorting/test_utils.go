package sorting

import "testing"

func runTests(t *testing.T, tests []struct {
	input    []int
	expected []int
}, sorting Sorting) {
	for _, test := range tests {
		// given
		actual := make([]int, len(test.input))
		copy(actual, test.input)

		// when
		sorting.Sort(actual)

		// then
		if !compareArray(actual, test.expected) {
			t.Errorf("Test failed! Expected = %v, Actual = %v", test.expected, actual)
		}
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
