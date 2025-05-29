package removeduplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expected_array []int
	}{
		{"Case 1", []int{1, 1, 2}, []int{1, 2}},
		{"Case 2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := removeDuplicates(tc.input)
			if length := len(tc.expected_array); output != length {
				t.Errorf("output: %d; expected: %d", output, length)
			}
			if !t.Failed() {
				for i, e := range tc.expected_array {
					if e != tc.expected_array[i] {
						t.Errorf("array[%d]= %d; expected: %d", i, tc.expected_array[i], e)
					}
				}
			}
		})
	}

}
