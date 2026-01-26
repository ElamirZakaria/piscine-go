package piscine

func Capitalize(s string) string {
	var b int = 0
	var str string = ""
	for _, a := range s {
		if (a >= 'a' && a <= 'z') && b == 0 {
			str = str + string(a-32)
			b = 1
			continue
		}
		if (a >= 'A' && a <= 'Z') && b == 0 {
			str = str + string(a)
			b = 1
			continue
		}
		if !(a >= 'a' && a <= 'z') {
			b = 0
		}
		str = str + string(a)
	}
	return str
}
