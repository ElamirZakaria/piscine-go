package piscine

func Compact(ptr *[]string) int {
	str := make([]string, 0)
	count := 0
	for _, a := range *ptr {
		if a != "" {
			str = append(str, a)
			count++
		}
	}
	*ptr = str
	return count
}
