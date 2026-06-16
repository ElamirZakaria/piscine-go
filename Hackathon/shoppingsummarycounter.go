package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	shoppingSummary := make(map[string]int)

	sliceOfStrings := make([]string, 0)

	s := ""

	for _, a := range str {
		if a == ' ' {
			sliceOfStrings = append(sliceOfStrings, s)
			s = ""
		} else {
			s += string(a)
		}
	}
	sliceOfStrings = append(sliceOfStrings, s)
	for _, a := range sliceOfStrings {
		shoppingSummary[a]++
	}
	return shoppingSummary
}
