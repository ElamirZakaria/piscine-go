package piscine

func MakeRange(min, max int) []int {
	size := max - min
	var answer []int
	if size <= 0 {
		answer = make([]int, 0)
	} else {
		answer = make([]int, size)
		for i := 0; i < size; i++ {
			answer[i] = min + i
		}
	}
	return answer
}
