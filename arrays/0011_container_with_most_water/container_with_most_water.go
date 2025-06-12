package containerwithmostwater


func maxArea(height []int) int {
	maxWater := 0
    left, right := 0, len(height)-1

    for left < right {
        water := min(height[left], height[right]) * (right - left)
        maxWater = max(maxWater, water)

        if height[left] < height[right] {
            left++
        } else if height[left] > height[right] {
            right--
        } else {
            left++
            right--
        }
    }
    return maxWater

}







func maxArea1(height []int) int {
	length := len(height)
	r := make(map[int]struct{},length)
	var val, sum  int 
	val = 0
	for i := length-1;  i > 0;  i-- {
		if height[i] > val {
			val = height[i]
			r[i] = struct{}{} 
		}
	}
	val = 0
	for i, h := range height {
		if h > val {
			val = h
			for j := range r {
				if n:= (j-i)*min(h,height[j]);  n > sum {
					sum = n
				}
			}
		}
	}
	return sum
}

























func maxArea_1_0(height []int) int {
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
func maxArea_1_1(height []int) int {
	r := 0
	l := len(height) -1
	
	var mi, d, s, vx1, vx2 int
		for _, v := range height {
			if v > vx1 {
			vx2 = vx1
				vx1 = v

				
			}else if v > vx2 {
			vx2 = v
				
			}
			
		}


	for {
		mi = min(height[r],height[l])
d = l-r
		s = mi * d
		if vx1 * (d-1)<= s   {
			if vx2 != mi {
				
				return s
			}
			if vx2 >= d {
				return s
			}
		}
		if height[r] < height[l] {
			for i := r + 1;  ; i++ {
				if vx2 * (l - i)  <=s {
					return s
				}
				if height[r] < height[i]  {
					if vx2 * (l - i) > s {
						r = i
						break
					}
				}
			}
			
		}else if height[r] > height[l] {
			for j := l - 1;  ; j-- {
				if vx2 * (j - r) <= s{
					return s
				}
				if height[l] < height[j]  {
					if vx2 *  (j-r) > s {
						l = j
						break
					}
				}
			}
		}else{
			s2 := maxArea_1_1(height[r+1:l])
			if s2>s {
				return s2
			}
		return s
		}
	}
}


