package piscine

func BasicAtoi(s string) int {
	var result int = 0
	temp := 0
	for i := 0; i < len(s); i++ {
		temp = int(s[i] - '0')
		result = result*10 + temp
	}
	return result
}
