package plusone

import (
	"slices"
	"testing"
)




func TestPlusOne(t *testing.T) {
	tests := []struct{
		name string
		input, expected []int
	}{
		{"Case 1",[]int{1,2,3},[]int{1,2,4}},
		{"Case 2",[]int{4,3,2,1},[]int{4,3,2,2}},
		{"Case 3",[]int{9},[]int{1,0}},
		{"Case 4",[]int{0},[]int{1}},
		{"Case 4",[]int{9,9,9},[]int{1,0,0,0}},
	}
	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			output := plusOne(tc.input)
			if !slices.Equal(tc.expected,output) {
				t.Errorf("Error, expected: %d, output: %d",tc.expected,output)
			}
		})
	}
	
}

