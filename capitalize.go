package piscine

func Capitalize(s string) string {
	var b int = 0
	var str string = ""
	for _, a := range s {
		if IsAlpha(string(a)) {
			if IsNumeric(string(a)) {
				str = str + string(a)
				b = 1
			} else if (IsLower(string(a))) && (b == 0) {
				str = str + string(a-32)
				b = 1
			} else if (IsUpper(string(a))) && (b == 0) {
				str = str + string(a)
				b = 1
			} else if (IsLower(string(a))) && (b == 1) {
				str = str + string(a)
				b = 1
			} else if (IsUpper(string(a))) && (b == 1) {
				str = str + string(a+32)
				b = 1
			}
		} else {
			str = str + string(a)
			b = 0
		}
	}
	return str
}
