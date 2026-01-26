package piscine

// import "piscine"

func IsAlpha(s string) bool {
	var ans bool = true

	for _, a := range s {
		if (a >= '0' && a <= '9') || IsUpper(string(a)) || IsLower(string(a)) {
		} else {
			ans = false
			break
		}
	}
	return ans
}
