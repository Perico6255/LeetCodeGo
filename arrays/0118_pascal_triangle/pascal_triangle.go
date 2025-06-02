package pascaltriangle

func generate(numRows int) [][]int {
	p := make([][]int,numRows)
	for i := range numRows {
		p[i]= make([]int, i + 1)
		p[i][0] = 1
		p[i][i] = 1
	}

	for i, p1 := range p {
		for j, p2 := range p1 {
			if p2 == 0 {
				p1[j] = p[i-1][j-1] + p[i-1][j]
			}
		}
	}
    
	return p 
}
