package piscine

func StringToIntSlice(str string) []int {
	var sti []int
	for _, a := range str {
		sti = append(sti, int(a))
	}
	return sti
}
