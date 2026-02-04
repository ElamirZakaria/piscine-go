package piscine

func Atoi(s string) int {
	var result int = 0
	temp := 0
	sign := 1
	flag := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			flag++
			temp = int(s[i] - '0')
			result = result*10 + temp
		} else if s[i] == ' ' {
			continue
		} else if s[i] == '-' {
			flag++
			if flag > 1 {
				result = 0
				break
			}
			sign = -1

		} else if s[i] == '+' {
			flag++
			if flag > 1 {
				result = 0
				break
			}
			sign = 1

		} else {
			result = 0
			break
		}
	}
	return result * sign
}
