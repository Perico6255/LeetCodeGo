package mergesortedarray



func merge(nums1 []int, m int, nums2 []int, n int)  {
	m--
	n--
	for i:= m + n +1  ; m>=0 && n >= 0 ; i--{
		if nums1[m]> nums2[n] {
			nums1[i]= nums1[m]
			m--
		}else{
			nums1[i]=nums2[n]
			n--
		}
	}	
	for ; n>=0; n--{
		nums1[n]= nums2[n]
		
	}
}
