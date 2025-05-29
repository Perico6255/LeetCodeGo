package lengthoflastword


func lengthOfLastWord(s string) int {
	l := len(s)-1
	n   := 0
	o   := 0
	for i := l; s[i] ==' ' ; i-- {
		n++
	}
	for i := l - n;s[i]>=0; i-- {
		if s[i] ==' ' {
			break
		}
		o++
	}
	return o
}
