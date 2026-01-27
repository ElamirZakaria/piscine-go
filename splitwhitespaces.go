package piscine

func SplitWhiteSpaces(s string) []string {
	var answer []string
	str := " "
	for _, a := range s {
		if (a == ' ') || (a == '\n') || (a == '\t') {
			answer = append(answer, str)
			str = ""

		} else {
			str = str + string(a)
		}
	}
	answer = append(answer, str)
	return answer
}
