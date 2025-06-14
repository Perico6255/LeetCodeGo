package sum

func threeSum(nums []int) [][]int {
	l := len(nums)
	sum := make([][]int,0)
	mi := make(map[int]struct{})
	mj := make(map[int]struct{})
	var n, nk   int
	for i := range l - 2  {
		if _,ok := mi[nums[i]]; ok {
			continue
		}
		if _,ok := mj[nums[i]]; ok {
			continue
		}
		for j:= i + 1; j < l - 1; j++ {
			if _,ok := mj[nums[j]]; ok {
				continue
			}
			if nk !=0 && nums[nk] == nums[j] {
				continue
				
			}
			n = nums[j] + nums[i]
			for k := j + 1; k< l; k++ {
				if -n == nums[k] {
					mi[nums[i]] = struct{}{}
					mj[nums[j]] = struct{}{}
					sum = append(sum, []int{nums[i],nums[j],nums[k]})
					nk = k
					break
				}
			}
		}
		nk = 0
		
	}
	return sum 
}
