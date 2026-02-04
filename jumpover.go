package piscine

func JumpOver(str string) string {
	s := ""
	if len(str) < 3 {
		s = "\n"
	} else {
		for i := 2; i < len(str); i += 3 {
			s = s + string(str[i])
		}
		s = s + "\n"
	}
	return s
}
