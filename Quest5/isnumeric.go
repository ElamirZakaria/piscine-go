package piscine

func IsNumeric(s string) bool {
	var ans bool = true

	for _, a := range s {
		if a >= '0' && a <= '9' {
		} else {
			ans = false
			break
		}
	}
	return ans
}
