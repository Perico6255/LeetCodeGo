package permutations

import "testing"


func TestPermute(t *testing.T) {
	tests := []struct{
		name string
		nums []int
		expected [][]int
	}{
		{"Case 1", []int{1,2,3}, [][]int{{1,2,3},{1,3,2},{2,1,3},{3,1,2},{3,2,1}}},
		{"Case 2", []int{0,1}, [][]int{{0,1},{1,0}}},
		{"Case 3", []int{2,3,4,5}, [][]int{{2,3,4,5},{2,3,5,4},{2,4,3,5},{2,4,5,3},{2,5,3,4},{2,5,4,3},{3,2,4,5},{3,2,5,4},{3,4,2,5},{3,4,5,2},{3,5,2,4},{3,5,4,2},{4,2,3,5},{4,2,5,3},{4,3,2,5},{4,3,5,2},{4,5,2,3},{4,5,3,2},{5,2,3,4},{5,2,4,3},{5,3,2,4},{5,3,4,2},{5,4,2,3},{5,4,3,2},},
	},}
	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			got := permute(tc.nums)
			if len(tc.expected) != len(got) {
				t.Fatalf("Not even same lenght, expected: %d, got: %d", tc.expected, got)
			}
		})
		
	}

	
}
func TestPermute2(t *testing.T) {
	tests := []struct{
		name string
		nums []int
		expected [][]int
	}{
		{"Case 1", []int{1,2,3}, [][]int{{1,2,3},{1,3,2},{2,1,3},{3,1,2},{3,2,1}}},
		{"Case 1", []int{1,2,3,4,5}, [][]int{}},
	}
	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			got := permute(tc.nums)
			if len(tc.expected) != len(got) {
				t.Fatalf("Not even same lenght, expected: %d, got: %d", tc.expected, got)
			}
		})
		
	}

	
}
