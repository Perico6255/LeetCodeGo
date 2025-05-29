package areoccurrencesequal

func areOccurrencesEqual(s string) bool {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	r := rune(s[0])
	n := m[r]
	for _, i := range m {
		if i != n {
			return false
		}
	}

	return true
}
