package piscine

func Sqrt(nb int) int {
	var i int
	result := 0
	for i = 0; i <= nb; i++ {
		if i*i == nb {
			result = i
			break
		} else if i*i > nb {
			result = 0
			break
		}
	}
	return result
}
