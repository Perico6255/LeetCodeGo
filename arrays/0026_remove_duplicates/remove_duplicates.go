package removeduplicates

func removeDuplicates(nums []int) int {
	c := 0
	a := -101
	for _, n := range nums {
		if n > a {
			a, nums[c] = n, n
			c++
		}

	}
	return c
}
