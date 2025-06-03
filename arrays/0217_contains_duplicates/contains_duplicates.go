package containsduplicates

func containsDuplicate(nums []int) bool {

	l := len(nums)
	m := make(map[int]struct{},l)
	ok := false
	for i := 0; i<l && !ok; i++ {
			_, ok = m[nums[i]]
			m[nums[i]] = struct{}{}
	}
    
	return ok
}
