package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	length := 0
	for range runes {
		length++
	}
	if n > length {
		return 0
	} else if n <= 0 {
		return 0
	} else {
		return runes[n-1]
	}
}
