package piscine

func SplitWhiteSpaces(s string) []string {
	var answer []string
	str := ""
	flag := 0
	for _, a := range s {
		if ((a == ' ') || (a == '\n') || (a == '\t')) && (flag == 1) {
			answer = append(answer, str)
			str = ""
			flag = 0

		} else {
			str = str + string(a)
			flag = 1
		}
	}
	// answer = append(answer, str)
	return answer
}
