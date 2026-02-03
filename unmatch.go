package piscine

func Unmatch(a []int) int {
	max := 0
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	count := make([]int, max+1)

	for _, i := range a {
		count[i]++
	}
	for j, i := range count {
		if i == 1 {
			return j
		}
	}
	return -1
}
