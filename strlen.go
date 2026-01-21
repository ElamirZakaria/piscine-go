package piscine

func Strlen(s string) int {
	length := 0
	for range s {
		length++
	}

	return length
}
