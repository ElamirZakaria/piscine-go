package piscine

func IterativeFactorial(nb int) int {
	var result int = 1
	for i := 2; i <= nb; i++ {
		result = result * i
	}
	return result
}
