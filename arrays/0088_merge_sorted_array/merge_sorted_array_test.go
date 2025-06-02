package mergesortedarray

import (
	"slices"
	"testing"
)



func TestMerge(t *testing.T) {
	tests := []struct{
		name string
		nums1, nums2, expected []int
		m, n int


	}{
		{"Case 1", []int{1,2,3,0,0,0},[]int{2,5,6},[]int{1,2,2,3,5,6},3,3},
		{"Case 2", []int{1},[]int{},[]int{1},1,0},
		{"Case 3", []int{0},[]int{1},[]int{1},0,1},
		{"Case 4", []int{3,0,0,0},[]int{1,2,4},[]int{1,2,3,4},1,3},
	}
	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			merge(tc.nums1,tc.m,tc.nums2,tc.n)
			if !slices.Equal(tc.expected,tc.nums1) {
				t.Errorf("Error, expected: %d, output: %d",tc.expected,tc.nums1)
			}
			
		})
	}
}
