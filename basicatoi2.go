package piscine

func BasicAtoi2(s string) int {
	var result int = 0
	temp := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			temp = int(s[i] - '0')
			result = result*10 + temp
		} else {
			continue
		}
	}
	return result
}
