package piscine

func Atoi(s string) int {
	var result int = 0
	temp := 0
	sign := 1
	negative := 0
	postive := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			temp = int(s[i] - '0')
			result = result*10 + temp
		} else if s[i] == ' ' {
			continue
		} else if s[i] == '-' {
			negative++
			sign = -1
			if negative == 2 {
				result = 0
				break
			}
		} else if s[i] == '+' {
			sign = 1
			postive++
			if postive == 2 {
				result = 0
				break
			}
		} else {
			result = 0
			break
		}
	}
	return result * sign
}
