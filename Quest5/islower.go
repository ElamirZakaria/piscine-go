package piscine

func IsLower(s string) bool {
	var ans bool = true

	for _, a := range s {
		if a >= 'a' && a <= 'z' {
		} else {
			ans = false
			break
		}
	}
	return ans
}
