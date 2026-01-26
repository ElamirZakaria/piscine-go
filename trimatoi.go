package piscine

func TrimAtoi(s string) int {
	// temp := 0
	result := 0
	sign := 1
	for _, a := range s {
		if result == 0 && a == '-' {
			sign = -1
		} else {
			if a >= '0' && a <= '9' {
				result = result*10 + int(a-'0')
			}
		}
	}
	return result * sign
}
