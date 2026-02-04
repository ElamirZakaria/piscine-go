package piscine

func LoafOfBread(str string) string {
	if len(str) <= 5 {
		return "Invalid Output"
	}
	s := ""
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		} else {
			s += string(str[i])
		}
		if (len(s)-count)%5 == 0 && len(s) != 0 {

			s = s + " "
			count++
			i++

		}
	}
	if s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s + "\n"
}
