package removeelemnt


func removeElement(nums []int, val int) int {
	l := len(nums)
	for i := range l {
		if nums[i] == val {
			for l> i{
				l--
				nl := nums[l]
				if nl != val {
					nums[i]=nl
					break
				}
			}
		}
	}
	return l 
}
