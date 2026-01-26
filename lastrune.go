package piscine

func LastRune(s string) rune {
	var length int = 0

	runes := []rune(s)
	for range s {
		length++
	}

	return runes[length-1]
}
