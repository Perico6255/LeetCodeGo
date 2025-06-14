package sum

func threeSum(nums []int) [][]int {
	m := make(map[int]int)
	r := make(map[int]struct{},len(nums))
	l := 0
	for _, n := range nums {
		if _,ok := m[n]; !ok {
			nums[l] = n
			l++
		}else {
			r[n] = struct{}{}
		}
		m[n]++
	}
	sum := make([][]int,0)



	for vr:= range r {

		l = -2*vr
		if l == 0 {
			continue
		}
		if  _,ok:= m[l]; ok  {
			sum = append(sum, []int{vr,vr,l})
			break
		}
	}
	if m[0] >2 {
		sum = append(sum, []int{0,0 , 0})
	}


	return sum 
}
