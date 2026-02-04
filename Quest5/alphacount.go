package piscine

func AlphaCount(s string) int {
	count := 0
	// characters := []byte(s)
	for _, a := range s {
		if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') {
			count++
		} else {
			continue
		}
	}
	return count
}
