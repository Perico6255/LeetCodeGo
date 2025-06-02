package plusone


func plusOne(digits []int) []int {
	l := len(digits)

	for i := l-1; i >0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] < 10 {
			return digits
		}
		digits[i]=0
	}

	if digits[0]==9 {
		s := make([]int, l+1)
		s[0] = 1
		s[1] = 0

		for i := 2; i < l ; i++ {
			s[i] = digits[i-1]
		}
		return s
		
	}
	digits[0]++
	return digits
    
}
