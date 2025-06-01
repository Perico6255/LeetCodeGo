package removeelemnt

import "testing"


func TestRemoveElemnt(t *testing.T) {
	tests := []struct{
		name string
		nums  []int
		val int
		expectedNums []int
		
	}{
		{"Case 1",[]int{3,2,2,3},3,[]int{2,2}},
		{"Case 2",[]int{0,1,2,2,3,0,4,2},2,[]int{0,1,4,0,3}},
		{"Case 3",[]int{1},1,[]int{}},
		{"Case 4",[]int{4,5},5,[]int{4}},
		{"Case 5",[]int{},0,[]int{}},
	}
	for _,tc := range tests {
		t.Run(tc.name,func(t *testing.T) {
			output := removeElement(tc.nums,tc.val)
			if length := len(tc.expectedNums); output != length  {
				t.Errorf("output: %d; expected: %d", output, length)
			}
			if !t.Failed() {
				for i, e := range tc.expectedNums {
					if e != tc.expectedNums[i] {
						t.Errorf("array[%d]= %d; expected: %d", i, tc.expectedNums[i], e)
					}
				}
			}
		})
		
	}
	
}
