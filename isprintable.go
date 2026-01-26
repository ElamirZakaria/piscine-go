package piscine

func IsPrintable(s string) bool {
	var ans bool = true

	for _, a := range s {
		if a < 32 || a == 127 {
			ans = false
			break
		} else {
		}
	}
	return ans
}
