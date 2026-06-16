package piscine

func Rot14(s string) string {
	str := ""
	for _, a := range s {
		if a >= 'a' && a <= 'z' {

			temp := a + 14

			if temp > 'z' {
				temp -= 26
			}
			str += string(temp)
		} else if a >= 'A' && a <= 'Z' {
			temp := a + 14
			if temp > 'Z' {
				temp -= 26
			}
			str += string(temp)
		} else {
			str += string(a)
		}
	}
	return str
}
