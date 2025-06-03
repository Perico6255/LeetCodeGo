package containsduplicates

import "testing"



func TestContainsDuplicates(t *testing.T) {
	tests := []struct{
		name string
		nums []int
		expected bool
	}{
		{"Case 1",[]int{1,2,3,1},true},
		{"Case 2",[]int{1,2,3,4},false},
		{"Case 3",[]int{1,1,1,3,3,4,3,2,4,2},true},
	}

	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			got := containsDuplicate(tc.nums)
			if got != tc.expected {
				t.Errorf("expected: %t, got: %t",tc.expected,got)
			}
		})
		
	}
}
