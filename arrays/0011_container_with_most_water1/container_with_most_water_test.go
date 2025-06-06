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






