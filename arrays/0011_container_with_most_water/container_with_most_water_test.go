package containerwithmostwater


import (
	"testing"
)


func TestMaxArea(t *testing.T) {

	tests := []struct{
		name string
		height []int
		expected int
	}{
		{"Case 1", []int{1,8,6,2,5,4,8,3,7}, 49},
		{"Case 2", []int{1,1}, 1},
		{"Case 3", []int{1,5,8,2,5,4,8,8,5}, 40},
		{"Case 4", []int{8,7,2,1}, 7},
		{"Case 5", []int{2,3,4,5,18,17,6}, 17},
	}
	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			out := maxArea(tc.height)
			if out != tc.expected {
				t.Errorf("Expected: %d, output: %d",tc.expected, out)
				
			}
		
		})
	}
	
}






