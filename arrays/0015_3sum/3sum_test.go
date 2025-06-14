package sum

import (
	"slices"
	"testing"
)


func TestThreeSum(t *testing.T) {
	tests := []struct{
		name string
		nums []int
		expected [][]int
	}{
		{"Case 1",[]int{-1,0,1,2,-1,-4},[][]int{{-1,-1,2},{-1,0,1}}},
	}
	
	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			got := threeSum(tc.nums)
			if len(got) != len(tc.expected) {
				t.Fatalf("Not same lenght \nexpected: %d, got: %d",tc.expected, got)
			}	
			for _, e := range tc.expected {
				for _, g := range got {
					if slices.Equal(e,g) {
						break
					}
				}
				t.Errorf("Expected array: %d not found in the output: %d",e, got)
			}
		})
		
	}
}
