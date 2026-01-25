package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < 2147483647; i++ {
		if IsPrime(i) == true {
			return i
		} else {
		}
	}
	return nb
}
