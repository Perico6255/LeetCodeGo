package permutations


func permute(nums []int) [][]int {
	l := len(nums)
	// if l == 1 {
	// 	return [][]int{nums}
	// }
	var c = 1
	for i := 2; i <= l; i++ {
		c *= i
	}
	sol := make([][]int, c)
	c = 0
	step(sol,nums,&c,l,l)



	return sol
}

func swap(nums []int, l, n1, n2 int) []int {
	n := make([]int,l)
	copy(n,nums)
	n[n1] = nums[n2]
	n[n2] = nums[n1]
	return n
}


func step(sol [][]int, n0 []int, s *int, p, l int)  {
	if p == 2 {
		sol[*s] = n0
		*s++
		sol[*s] = swap(n0, l, l-1, l-2)
		*s++
		return
	}
	step(sol,n0,s,p-1,l)
	for i := 1; i < p; i++ {
		step(sol, 
			swap(n0,l, l- p, l-i ),
			s,
			p - 1,
			l)
	}
}




