package piscine

func ReverseMenuIndex(menu []string) []string {
	reverseMenue := make([]string, len(menu))
	j := 0
	for i := len(menu) - 1; i >= 0; i-- {
		reverseMenue[j] = menu[i]
		j++
	}
	return reverseMenue
}
