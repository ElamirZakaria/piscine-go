package piscine

func Join(strs []string, sep string) string {
	str := ""
	length := 0
	for range strs {
		length++
	}
	for i, a := range strs {
		if i != length {
			str = str + a + sep
		} else {
			str += a
		}
	}
	return str
}
