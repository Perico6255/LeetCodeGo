package sum

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return [][]int{}
	}
	slices.Sort(nums)
	t := make([][]int, 0)
	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		low := i + 1
		high := l - 1
		var s int
		for low < high {
			s = v + nums[low] + nums[high]
			switch {
			case s < 0:
				low++
			case s > 0:
				high--
			case s == 0:
				last_low := nums[low]
				last_high := nums[high]

				t = append(t, []int{v, last_low, last_high})
				for low < high && nums[low] == last_low {
					low++
				}
				for low < high && nums[high] == last_high {
					high--
				}
			}
		}
	}
	return t
}

// func threeSum(nums []int) [][]int {
// 	m := make(map[int]int)
// 	r := make(map[int]struct{},len(nums))
// 	l := 0
// 	for _, n := range nums {
// 		if _,ok := m[n]; !ok {
// 			nums[l] = n
// 			l++
// 		}else {
// 			r[n] = struct{}{}
// 		}
// 		m[n]++
// 	}
// 	sum := make([][]int,0)
//
//
//
// 	for vr:= range r {
//
// 		l = -2*vr
// 		if l == 0 {
// 			continue
// 		}
// 		if  _,ok:= m[l]; ok  {
// 			sum = append(sum, []int{vr,vr,l})
// 			break
// 		}
// 	}
// 	if m[0] >2 {
// 		sum = append(sum, []int{0,0 , 0})
// 	}
//
//
// 	return sum
// }
