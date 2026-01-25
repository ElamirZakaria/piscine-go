package piscine

func RecursiveFactorial(nb int) int {
	if nb <= 1 || nb > 20 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
