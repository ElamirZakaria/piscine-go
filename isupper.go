package piscine

func IsUpper(s string) bool {
	var ans bool = true

	for _, a := range s {
		if a >= 'A' && a <= 'Z' {
		} else {
			ans = false
			break
		}
	}
	return ans
}
