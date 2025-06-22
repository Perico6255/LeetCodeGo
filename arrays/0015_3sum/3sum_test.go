package sum

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{"Case 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"Case 2", []int{0, 1, 1}, [][]int{}},
		{"Case 3", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := threeSum(tc.nums)
			if len(got) != len(tc.expected) {
				t.Fatalf("Not same lenght \nexpected: %d, got: %d", tc.expected, got)
			} else {
				t.Log("Same Lenght!!")
				t.Logf("expected: %d", tc.expected)
				t.Logf("got: %d", got)
			}

		})

	}
}
