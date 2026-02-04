package piscine

func AppendRange(min, max int) []int {
	var answer []int
	j := max - min

	for i := 0; i < j; i++ {
		answer = append(answer, min+i)
	}
	return answer
}
