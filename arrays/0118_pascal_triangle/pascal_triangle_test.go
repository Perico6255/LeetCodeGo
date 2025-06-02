package pascaltriangle

import (
	"slices"
	"testing"
)


func TestGenerate(t *testing.T) {
	tests := []struct{
		name string
		input int
		expected [][]int
	}{
		{"Case 1",1, [][]int{
			{1},
		}},
		{"Case 2",2, [][]int{
			{1},
			{1,1},
		}},
		{"Case 3",3, [][]int{
			{1},
			{1,1},
			{1,2,1},
		}},
		{"Case 4",4, [][]int{
			{1},
			{1,1},
			{1,2,1},
			{1,3,3,1},
		}},
		{"Case 5",5, [][]int{
			{1},
			{1,1},
			{1,2,1},
			{1,3,3,1},
			{1,4,6,4,1},
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			out := generate(tc.input)
			for i := range tc.input {
				if !slices.Equal(out[i],tc.expected[i]) {
					t.Errorf("line: %d, expected: %d, got: %d",i,tc.expected[i],out[i])
				}
				
			}
		})
		
	}


	
}
