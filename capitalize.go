package piscine

func IsChar(s rune) bool {
	if IsLower(string(s)) || IsUpper(string(s)) {
		return true
	}
	return false
}

func Capitalize(s string) string {
	var b int = 0
	var str string = ""
	for i, a := range s {
		if (IsLower(string(a))) && b == 0 && IsLower(string(s[i+1])) {
			str = str + string(a-32)
			b = 1
			//	continue
		} else if (IsUpper(string(a))) && b == 0 {
			str = str + string(a)
			b = 1
			// continue
		} else if !((IsLower(string(a))) || (IsUpper(string(a)))) {
			b = 0
			str = str + string(a)
		} else {
			if IsUpper(string(a)) {
				str = str + string(a+32)
			} else {
				str = str + string(a)
			}
		}
	}
	return str
}

// Capitalize("6n\\#l&RxHNXc9") == "6N\\#L&RxHNXC9" instead of "6n\\#L&Rxhnxc9"  6n\#l&Rxhnxc9
