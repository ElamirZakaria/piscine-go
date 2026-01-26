package piscine

func Index(s string, toFind string) int {
	runes := []rune(toFind)
	lengths := 0
	for range runes {
		lengths++
	}
	if lengths == 0 {
		return -1
	}
	for i, a := range s {
		if a == runes[0] {
			return i
		} else {
			continue
		}
	}
	return -1
}
