package piscine

func Index(s string, toFind string) int {
	runes := []rune(toFind)
	for i, a := range s {
		if a == runes[0] {
			return i
		} else {
			continue
		}
	}
	return -1
}
