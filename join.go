package piscine

func Join(strs []string, sep string) string {
	str := ""
	for _, a := range strs {
		str = str + a + sep
	}
	return str
}
