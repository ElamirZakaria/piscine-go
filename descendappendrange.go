package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}
	a := []int{}

	for i := max; i > min; i-- {
		a = append(a, i)
	}
	return a
}
