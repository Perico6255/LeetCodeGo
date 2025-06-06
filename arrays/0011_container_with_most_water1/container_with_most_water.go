package containerwithmostwater

func maxArea(height []int) int {
	var v int
	var mi int
	for i := 0; i < len(height)-1; i++ {
		for j := i+1; j < len(height); j++ {
			mi = min(height[i],height[j])
			if a := mi * (j-i);  v <a  {
				v = a
			}
			
		}
		
	}

	return v
}


